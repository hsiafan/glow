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
		err = cmd.ParseOsArgsAndExecute()
		if err != nil {
			fmt.Print(err)
		}

	3. For composite command line:

		cc := flagx.NewCompositeCommand("my_cc", "some description")
		option := &Option{}
		_ = cc.AddSubCommand("my_command", "some description", option, func() error {
			return myHandle(option)
		})
		if err := cc.ParseAndExecute(os.Args[1:]); err != nil {
			if err == flag.ErrHelp {

			} else {
				fmt.Println(err)
			}
		}

	4. struct field tag:
		name:			The arg name. if not set, use converted struct filed name
		default:		Default arg value
		description:	Arg usage and other messages
		args:			ture|false. Mark this field as positional arg, name will be ignored, index works.
		index:			The arg index. works when args is true
		ignore:			Ignore this field, do not parse and add arg flag

	5. struct field type:
		string
		bool
		int
		int64
		uint
		uint64
		float64
		time.Duration
	Pointer and slice type are not supported. However, positional args do support slices type to handle var len args.
*/
package flagx
