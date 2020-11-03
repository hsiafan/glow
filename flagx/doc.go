/*
Package flagx provide convenient for parsing command line arguments, and execute logics.

Usage:

1. Define one Option struct:

	type Option struct {
		Address string `description:"the address"`
		Path    string `description:"the path"`
	}

2. For plain command line:

	option := &Option{}
	cmd, err := flagx.NewCommand("my_command", "some description", option, func() error {
		return myHandle(option)
	})
	if err != nil {
		fmt.Print("parse arguments failed", err)
		return
	}
	cmd.ParseOsArgsAndExecute()

3. For composite command line:

	cc := flagx.NewCompositeCommand("my_cc", "some description")
	option := &Option{}
	_ = cc.AddSubCommand("my_command", "some description", option, func() error {
		return myHandle(option)
	})
	cc.ParseAndExecute(os.Args[1:])

4. struct field tag:
	name:			The arg name. if not set, use converted struct filed name
	default:		Default arg value
	description:	Arg usage and other messages
	"flag"          If is a flag arg(true) or a non-flag arg(false), default is true.
	index:			The position of non-flag args, start from 0. If a field with slice type, set args true, and leave index not set, will receive all non-consumed non-flag args.
	ignore:			Ignore this field, do not parse and add arg flag
	"required"      If this args is required(true|false). Default is false. If default value is set, the required filed will be ignored.

5. supported struct field type:
	string
	bool
	int
	int8
	int16
	int32
	int64
	uint
	uint8
	uint16
	uint32
	uint64
	float32
	float64
	time.Duration

Pointer and are not supported.
And, slice type with the element type above can be used, flag args can be set multi times, and stored in the slice, such as "-f 1 -f 2".
However, positional non-flag args can not use slices type.
*/
package flagx
