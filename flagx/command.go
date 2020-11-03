package flagx

import (
	"errors"
	"flag"
	"fmt"
	"github.com/hsiafan/glow/floatx"
	"github.com/hsiafan/glow/intx"
	"github.com/hsiafan/glow/reflectx"
	"github.com/hsiafan/glow/stringx"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	nameField         = "name"        // the name. if not set, use converted struct filed name
	defaultValueField = "default"     // the default value
	descriptionField  = "description" // the usage message
	flagField         = "flag"        // if is a flag arg(true) or a non-flag arg(false), default is true.
	indexField        = "index"       // the position of non-flag args, start from 0. If a field with slice type, set args true, and leave index not set, will receive all non-consumed non-flag args.
	ignoreFiled       = "ignore"      // for ignore one struct field
	requiredField     = "required"    // If this args is required(true|false). Default is false. If default value is set, the required filed will be ignored.
)

// alias command handle function
type Handle = func() error

// Command is a command line
type Command struct {
	Name             string                      // the name of this command
	Description      string                      // usage message
	parentCmd        string                      // the composite command name, if this is a sub command
	flagSet          *flag.FlagSet               // for internal process
	flagFields       []*fieldFlagValue           // flag field
	positionalFields map[int]*positionalArgField // for storing positional non-flag args field
	remainFields     *remainedArgsField          // for storing remained non-flag args
	handle           Handle
}

// NewCommand create new command
func NewCommand(Name string, Description string, option interface{}, handle Handle) (*Command, error) {
	flagSet := &flag.FlagSet{}

	v := reflect.ValueOf(option)
	if v.IsValid() == false {
		return nil, errors.New("not valid option value")
	}

	for v.Kind() == reflect.Ptr {
		if !v.IsNil() {
		} else {
			v.Set(reflect.New(v.Type().Elem()))
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, errors.New("option should be a struct")
	}

	var fieldFlagValues []*fieldFlagValue
	var positionalArgFields = map[int]*positionalArgField{}
	var raf *remainedArgsField

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := v.Type().Field(i)

		ignore, err := reflectx.GetBoolTagValue(fieldType.Tag, ignoreFiled, false)
		if err != nil {
			return nil, errors.New("invalid bool value for ignore tag of field:" + fieldType.Name)
		}
		if ignore {
			continue
		}

		if fieldValue.IsValid() == false || fieldValue.CanSet() == false {
			return nil, fmt.Errorf("invalid field %v", fieldType.Name)
		}

		//TODO: check field type

		var flagName string
		tagName := fieldType.Tag.Get(nameField)
		if tagName != "" {
			flagName = tagName
		} else {
			flagName = toFlagName(fieldType.Name)
		}

		isFlagArg, err := reflectx.GetBoolTagValue(fieldType.Tag, flagField, true)
		if err != nil {
			return nil, fmt.Errorf("struct tag of flag is not valid, field name: %v", fieldType.Name)
		}
		required, err := reflectx.GetBoolTagValue(fieldType.Tag, requiredField, false)
		if err != nil {
			return nil, fmt.Errorf("struct tag of required is not valid, field name: %v", fieldType.Name)
		}

		if !isFlagArg {
			sv, hasIndex := fieldType.Tag.Lookup(indexField)
			if hasIndex {
				index, err := strconv.Atoi(sv)
				if err != nil || index < 0 {
					return nil, fmt.Errorf("illegal index value %v for field %v", sv, fieldType.Name)
				}

				if _, ok := positionalArgFields[index]; ok {
					return nil, fmt.Errorf("field %v and %v have same index %v", positionalArgFields[index].fieldName, fieldType.Name, index)
				}
				if _, ok := positionalArgFields[index]; ok {
					if index == -1 {
						return nil, fmt.Errorf("both fields %v and %v want to receive all remains args", positionalArgFields[index].fieldName, fieldType.Name)
					} else {
						return nil, fmt.Errorf("both fields %v and %v have same index %v", positionalArgFields[index].fieldName, fieldType.Name, index)
					}
				}
				positionalArgFields[index] = &positionalArgField{
					value:     fieldValue,
					_type:     fieldType.Type,
					fieldName: fieldType.Name,
					name:      flagName,
					required:  required,
					index:     index,
				}
			} else {
				if fieldType.Type.Kind() != reflect.Slice {
					return nil, fmt.Errorf("field %v want to receive all remains args, but is not slice type", fieldType.Name)
				}
				if raf != nil {
					return nil, fmt.Errorf("both fields %v and %v want to receive remained non-flag args", raf.fieldName, fieldType.Name)
				} else {
					raf = &remainedArgsField{
						value:     fieldValue,
						_type:     fieldType.Type,
						fieldName: fieldType.Name,
						name:      flagName,
					}
				}
			}
			continue
		} else {
			defaultValue, hasDefault := fieldType.Tag.Lookup(defaultValueField)
			var ffv = &fieldFlagValue{
				name:         flagName,
				value:        fieldValue,
				defaultValue: defaultValue,
				_type:        fieldType.Type,
				required:     required,
			}
			fieldFlagValues = append(fieldFlagValues, ffv)
			description := fieldType.Tag.Get(descriptionField)
			if hasDefault {
				err := setValue(defaultValue, fieldType.Type.Kind(), fieldValue)
				if err != nil {
					return nil, fmt.Errorf("invalid default value for field %v, error: %w", fieldType.Name, err)
				}
			}
			flagSet.Var(ffv, flagName, description)
		}
	}

	// check if all non-flag args is consumed
	for idx := 0; idx < len(positionalArgFields); idx++ {
		if _, ok := positionalArgFields[idx]; !ok {
			return nil, fmt.Errorf("no field receive %v-th non-flag arg", idx)
		}
	}

	cmd := &Command{
		flagSet:          flagSet,
		flagFields:       fieldFlagValues,
		positionalFields: positionalArgFields,
		remainFields:     raf,
		Name:             Name,
		Description:      Description,
		handle:           handle,
	}

	flagSet.Usage = func() {
		output := flagSet.Output()
		if cmd.Description != "" {
			_, _ = fmt.Fprintln(output, cmd.Description+"\n")
		}

		argDes := argsDesc(raf, positionalArgFields)
		if cmd.parentCmd != "" {
			_, _ = fmt.Fprintf(output, "Usage: %s %s %s\n", cmd.parentCmd, cmd.Name, argDes)
		} else {
			_, _ = fmt.Fprintf(output, "Usage: %s %s\n", cmd.Name, argDes)
		}

		flagSet.PrintDefaults()
	}

	return cmd, nil

}

// ParseOsArgsAndExecute parse commandline passed arguments, and run handlers.If error occurred, will exit with non-zero code.
func (c *Command) ParseOsArgsAndExecute() {
	c.ParseAndExecute(os.Args[1:])
}

// ParseAndExecute parse arguments, and run handlers. If error occurred, will exit with non-zero code.
func (c *Command) ParseAndExecute(arguments []string) {
	if err := c.flagSet.Parse(arguments); err != nil {
		if err == flag.ErrHelp {
			// already show usage
			return
		}
		c.exitOnError(fmt.Errorf("parse flag error: %w", err))
		return
	}

	for _, field := range c.flagFields {
		if field.required && !field.set {
			c.exitOnError(fmt.Errorf("flag arg %v required but not set", field.name))
		}
	}

	for _, field := range c.positionalFields {
		if field.required && !field.set {
			c.exitOnError(fmt.Errorf("non-flag arg %v[%v] required but not set", field.name, field.index))
		}
	}

	// deal with positional args
	args := c.flagSet.Args()
	for i, arg := range args {
		if i >= len(c.positionalFields) {
			break
		}
		af := c.positionalFields[i]
		if err := af.Set(arg); err != nil {
			c.exitOnError(fmt.Errorf("set positional arguments error: %w", err))
			return
		}
	}

	if len(args) > len(c.positionalFields) {
		remainArgs := args[len(c.positionalFields):]
		if c.remainFields == nil {
			c.exitOnError(fmt.Errorf("still has args not handled: %v", remainArgs))
		}
		if err := c.remainFields.Set(remainArgs); err != nil {
			c.exitOnError(err)
			return
		}
	}

	if err := c.handle(); err != nil {
		c.exitOnError(err)
	}
}

// ShowUsage print formatted usage message
func (c *Command) ShowUsage() {
	c.flagSet.Usage()
}

func (c *Command) exitOnError(err error) {
	fmt.Println(err)
	c.ShowUsage()
	os.Exit(-1)
}

// make args description output for usage message
func argsDesc(raf *remainedArgsField, argFields map[int]*positionalArgField) string {
	var joiner = stringx.Joiner{
		Separator: " ",
	}

	for idx := 0; idx < len(argFields); idx++ {
		f := argFields[idx]
		joiner.Add(f.name)
	}

	if raf != nil {
		joiner.Add("[" + raf.name + "...]")
	}
	return joiner.String()
}

func toFlagName(filedName string) string {
	var sb strings.Builder

	for i := 0; i < len(filedName); i++ {
		c := filedName[i]
		if stringx.IsUpperASCII(c) {
			if sb.Len() != 0 {
				sb.WriteByte('-')
			}
			sb.WriteByte(stringx.ToLowerASCII(c))
		} else {
			sb.WriteByte(c)
		}
	}
	return sb.String()
}

var _ flag.Value = (*fieldFlagValue)(nil)

// flag.Value implementation that store value in a struct field
type fieldFlagValue struct {
	name         string
	defaultValue string
	_type        reflect.Type
	value        reflect.Value
	required     bool
	set          bool
}

func (f *fieldFlagValue) String() string {
	return f.defaultValue
}

func (f *fieldFlagValue) Set(s string) error {
	if f.set && f._type.Kind() != reflect.Slice {
		return errors.New("flag arg " + f.name + " already set")
	}
	f.set = true
	if f._type.Kind() != reflect.Slice {
		return setValue(s, f._type.Kind(), f.value)
	}
	newValue := reflect.New(f._type.Elem())
	if err := setValue(s, f._type.Elem().Kind(), newValue.Elem()); err != nil {
		return fmt.Errorf("set %v argument error: %w", f.name, err)
	}
	f.value.Set(reflect.Append(f.value, newValue.Elem()))
	return nil
}

func (f *fieldFlagValue) IsBoolFlag() bool {
	return f._type.Kind() == reflect.Bool
}

// for remember non-flag positional args.
type positionalArgField struct {
	value     reflect.Value
	_type     reflect.Type
	fieldName string // struct field name
	name      string // flag name
	index     int    // non-flag arg index
	required  bool
	set       bool
}

func (f *positionalArgField) Set(s string) error {
	f.set = true
	return setValue(s, f._type.Kind(), f.value)
}

// for remember non-flag remained args.
type remainedArgsField struct {
	value     reflect.Value
	_type     reflect.Type
	fieldName string // struct field name
	name      string // flag name
}

func (f *remainedArgsField) Set(args []string) error {
	slice := reflect.MakeSlice(f._type, len(args), len(args))
	for idx, arg := range args {
		if err := setValue(arg, f._type.Elem().Kind(), slice.Index(idx)); err != nil {
			return fmt.Errorf("set remained arguments error: %w", err)
		}
	}
	f.value.Set(slice)
	return nil
}

// set one arg field, the value converted from str input
func setValue(str string, kind reflect.Kind, value reflect.Value) error {
	switch kind {
	case reflect.String:
		value.SetString(str)
	case reflect.Int:
		v, err := intx.ParseInt(str)
		if err != nil {
			return err
		}
		value.SetInt(int64(v))
	case reflect.Int8:
		v, err := intx.ParseInt8(str)
		if err != nil {
			return err
		}
		value.SetInt(int64(v))
	case reflect.Int16:
		v, err := intx.ParseInt16(str)
		if err != nil {
			return err
		}
		value.SetInt(int64(v))
	case reflect.Int32:
		v, err := intx.ParseInt32(str)
		if err != nil {
			return err
		}
		value.SetInt(int64(v))
	case reflect.Int64:
		switch value.Interface().(type) {
		case time.Duration:
			v, err := time.ParseDuration(str)
			if err != nil {
				return err
			}
			value.Set(reflect.ValueOf(v))
			return nil
		}
		v, err := intx.ParseInt64(str)
		if err != nil {
			return err
		}
		value.SetInt(v)
	case reflect.Uint:
		v, err := intx.ParseUint(str)
		if err != nil {
			return err
		}
		value.SetUint(uint64(v))
	case reflect.Uint8:
		v, err := intx.ParseUint8(str)
		if err != nil {
			return err
		}
		value.SetUint(uint64(v))
	case reflect.Uint16:
		v, err := intx.ParseUint16(str)
		if err != nil {
			return err
		}
		value.SetUint(uint64(v))
	case reflect.Uint32:
		v, err := intx.ParseUint32(str)
		if err != nil {
			return err
		}
		value.SetUint(uint64(v))
	case reflect.Uint64:
		v, err := intx.ParseUint64(str)
		if err != nil {
			return err
		}
		value.SetUint(v)
	case reflect.Float32:
		v, err := floatx.Parse32(str)
		if err != nil {
			return err
		}
		value.SetFloat(float64(v))
	case reflect.Float64:
		v, err := floatx.Parse64(str)
		if err != nil {
			return err
		}
		value.SetFloat(v)
	case reflect.Bool:
		v, err := strconv.ParseBool(str)
		if err != nil {
			return err
		}
		value.SetBool(v)
	default:
		return fmt.Errorf("unsupported field type: %v", kind)
	}
	return nil
}
