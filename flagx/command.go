package flagx

import (
	"errors"
	"flag"
	"fmt"
	"github.com/hsiafan/glow/stringx"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/hsiafan/glow/floatx"
	"github.com/hsiafan/glow/intx"
	"github.com/hsiafan/glow/reflectx"
)

const (
	nameField         = "name"        // the name. if not set, use converted struct filed name
	defaultValueField = "default"     // the default value
	descriptionField  = "description" // the usage message
	argsField         = "args"        // mark field as args param, not option param
	indexField        = "index"       // the position of args, start from 0, used with argsField tag
	ignoreFiled       = "ignore"      // for ignore one struct field
)

// alias command handle function
type Handle = func() error

// Command is a command line
type Command struct {
	Name        string        // the name of this command
	Description string        // usage message
	parentCmd   string        // the composite command name, if this is a sub command
	flagSet     *flag.FlagSet // for internal process
	argFields   []argFiled    // for storing args field
	handle      Handle
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
	var argFields []argFiled

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

		var fieldName string
		tagName := fieldType.Tag.Get(nameField)
		if tagName != "" {
			fieldName = tagName
		} else {
			fieldName = toFlagName(fieldType.Name)
		}

		description := fieldType.Tag.Get(descriptionField)
		isArg, err := reflectx.GetBoolTagValue(fieldType.Tag, argsField, false)
		if err != nil {
			return nil, fmt.Errorf("struct tag of tags is not valid, field name: %v", fieldType.Name)
		}

		if isArg {
			index, err := reflectx.GetIntTagValue(fieldType.Tag, indexField, 0)
			if err != nil {
				return nil, fmt.Errorf("illegal index value for field %v", fieldType.Name)
			}
			f := argFiled{
				value: fieldValue,
				_type: fieldType.Type,
				name:  fieldName,
				index: index,
			}
			argFields = append(argFields, f)
			continue
		}

		switch fieldValue.Kind() {
		case reflect.Bool:
			v, err := reflectx.GetBoolTagValue(fieldType.Tag, defaultValueField, false)
			if err != nil {
				return nil, fmt.Errorf("invalid default value for field %v, error: %w", fieldType.Name, err)
			}
			flagSet.BoolVar((*bool)(unsafe.Pointer(fieldValue.Addr().Pointer())), fieldName, v, description)
		case reflect.Int:
			v, err := reflectx.GetIntTagValue(fieldType.Tag, defaultValueField, 0)
			if err != nil {
				return nil, fmt.Errorf("invalid default value for field %v, error: %w", fieldType.Name, err)
			}
			flagSet.IntVar((*int)(unsafe.Pointer(fieldValue.Addr().Pointer())), fieldName, v, description)
		case reflect.Int64:
			switch fieldValue.Interface().(type) {
			case time.Duration:
				v := reflectx.GetTagValue(fieldType.Tag, defaultValueField, "0")
				duration, err := time.ParseDuration(v)
				if err != nil {
					return nil, fmt.Errorf("invalid default value for field %v, error: %w", fieldType.Name, err)
				}
				flagSet.DurationVar((*time.Duration)(unsafe.Pointer(fieldValue.Addr().Pointer())), fieldName, duration, description)
			default:
				v, err := reflectx.GetInt64TagValue(fieldType.Tag, defaultValueField, 0)
				if err != nil {
					return nil, fmt.Errorf("invalid default value for field %v, error: %w", fieldType.Name, err)
				}
				flagSet.Int64Var((*int64)(unsafe.Pointer(fieldValue.Addr().Pointer())), fieldName, v, description)
			}

		case reflect.Uint:
			v, err := reflectx.GetUIntTagValue(fieldType.Tag, defaultValueField, 0)
			if err != nil {
				return nil, fmt.Errorf("invalid default value for field %v, error: %w", fieldType.Name, err)
			}
			flagSet.UintVar((*uint)(unsafe.Pointer(fieldValue.Addr().Pointer())), fieldName, v, description)
		case reflect.Uint64:
			v, err := reflectx.GetUInt64TagValue(fieldType.Tag, defaultValueField, 0)
			if err != nil {
				return nil, fmt.Errorf("invalid default value for field %v, error: %w", fieldType.Name, err)
			}
			flagSet.Uint64Var((*uint64)(unsafe.Pointer(fieldValue.Addr().Pointer())), fieldName, v, description)
		case reflect.Float64:
			v, err := reflectx.GetFloat64TagValue(fieldType.Tag, defaultValueField, 0)
			if err != nil {
				return nil, fmt.Errorf("invalid default value for field %v, error: %w", fieldType.Name, err)
			}
			flagSet.Float64Var((*float64)(unsafe.Pointer(fieldValue.Addr().Pointer())), fieldName, v, description)
		case reflect.String:
			v := reflectx.GetTagValue(fieldType.Tag, defaultValueField, "")
			flagSet.StringVar((*string)(unsafe.Pointer(fieldValue.Addr().Pointer())), fieldName, v, description)
		case reflect.Ptr:
			return nil, fmt.Errorf("field can not be pointer: %v", fieldType.Name)
		case reflect.Slice:
			return nil, fmt.Errorf("field can not be slice: %v", fieldType.Name)
		default:
			return nil, fmt.Errorf("unsupported field type: %v", fieldType.Name)
		}
	}

	cmd := &Command{
		flagSet:     flagSet,
		argFields:   argFields,
		Name:        Name,
		Description: Description,
		handle:      handle,
	}

	flagSet.Usage = func() {
		output := flagSet.Output()
		if cmd.Description != "" {
			_, _ = fmt.Fprintln(output, cmd.Description+"\n")
		}

		argDes := argsDesc(argFields)
		if cmd.parentCmd != "" {
			_, _ = fmt.Fprintf(output, "Usage: %s %s %s\n", cmd.parentCmd, cmd.Name, argDes)
		} else {
			_, _ = fmt.Fprintf(output, "Usage: %s %s\n", cmd.Name, argDes)
		}

		flagSet.PrintDefaults()
	}

	return cmd, nil

}

func argsDesc(argFields []argFiled) string {
	for _, af := range argFields {
		if af.value.Kind() == reflect.Slice {
			return af.name
		}
	}
	var sb strings.Builder
	sort.Slice(argFields, func(i, j int) bool {
		return argFields[i].index < argFields[j].index
	})
	for _, f := range argFields {
		if sb.Len() > 0 {
			sb.WriteRune(' ')
		}
		sb.WriteString(f.name)
	}
	return sb.String()
}

// ParseOsArgsAndExecute parse commandline passed arguments, and run handlers
func (c *Command) ParseOsArgsAndExecute() {
	c.ParseAndExecute(os.Args[1:])
}

// ParseAndExecute parse arguments, and run handlers
func (c *Command) ParseAndExecute(arguments []string) {
	if err := c.flagSet.Parse(arguments); err != nil {
		if err == flag.ErrHelp {
			// already show usage
			return
		}
		c.exitOnError(fmt.Errorf("parse flag error: %w", err))
		return
	}

	args := c.flagSet.Args()
	for _, af := range c.argFields {
		switch af.value.Kind() {
		case reflect.Slice:
			eType := af._type.Elem()
			slice := reflect.MakeSlice(af._type, len(args), len(args))
			//slice := slicePtr.Elem()

			for idx, arg := range args {
				if err := setField(arg, eType.Kind(), slice.Index(idx)); err != nil {
					c.exitOnError(fmt.Errorf("set positional arguments error: %w", err))
					return
				}
			}
			af.value.Set(slice)
		default:
			if af.index >= len(args) {
				c.exitOnError(fmt.Errorf("no enough args, require %v, but got: %v", af.index+1, len(args)))
				return
			}
			if err := setField(args[af.index], af.value.Kind(), af.value); err != nil {
				c.exitOnError(fmt.Errorf("set position args error, postion: %v, error: %w", af.index, err))
				return
			}
		}
	}

	if err := c.handle(); err != nil {
		c.exitOnError(err)
	}
}

func (c *Command) exitOnError(err error) {
	fmt.Println(err)
	c.ShowUsage()
	os.Exit(-1)
}

func setField(str string, kind reflect.Kind, value reflect.Value) error {
	switch kind {
	case reflect.String:
		value.SetString(str)
	case reflect.Int:
		v, err := intx.ParseInt(str)
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
	case reflect.Uint64:
		v, err := intx.ParseUint64(str)
		if err != nil {
			return err
		}
		value.SetUint(v)
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

// ShowUsage print formatted usage message
func (c *Command) ShowUsage() {
	c.flagSet.Usage()
}

type argFiled struct {
	value reflect.Value
	_type reflect.Type
	name  string
	index int
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
