package flagx

import (
	"errors"
	"fmt"
	"os"
)

// CompositeCommand
type CompositeCommand struct {
	Name        string       // the command name
	Description string       // the description
	subCommands []subCommand // sub commands
	hasHelp     bool         // if add help command
}

// Create new CompositeCommand
func NewCompositeCommand(Name string) *CompositeCommand {
	return &CompositeCommand{
		Name: Name,
	}
}

// Add one sub command
func (c *CompositeCommand) AddSubCommand(command *Command, handle func() error) {
	c.subCommands = append(c.subCommands, subCommand{command: command, handle: handle})
}

// Add a command which is name is help, and print usage
func (c *CompositeCommand) AddHelpCommand() {
	cmd, err := NewCommand("help", "show help message", &emptyStruct{})
	if err != nil {
		panic(err)
	}
	c.hasHelp = true
	c.AddSubCommand(cmd, func() error {
		c.ShowUsage()
		return nil
	})
}

// Parse commandline passed arguments, and execute command
func (c *CompositeCommand) ParseOsArgsAndExecute() error {
	return c.ParseAndExecute(os.Args[1:])
}

// Parse arguments, and execute command
func (c *CompositeCommand) ParseAndExecute(arguments []string) error {
	if len(arguments) == 0 {
		if c.hasHelp {
			arguments = []string{"help"}
		}
		return errors.New("should specify a sub command")
	}
	if len(arguments) == 1 && (arguments[0] == "-h" || arguments[0] == "-help") {
		arguments[0] = "help"
	}
	for _, sc := range c.subCommands {
		if sc.command.Name == arguments[0] {
			if err := sc.command.Parse(arguments[1:]); err != nil {
				return err
			}
			return sc.handle()
		}
	}
	return errors.New("unknown command: " + arguments[0])
}

// Show usage
func (c *CompositeCommand) ShowUsage() {
	if c.Description != "" {
		fmt.Println(c.Description + "\n")
	}
	fmt.Println("Usage:", c.Name)
	for _, sc := range c.subCommands {
		command := sc.command
		fmt.Println("  ", command.Name)
		fmt.Println("    ", command.Description)
	}
}

type subCommand struct {
	command *Command
	handle  func() error
}

type emptyStruct struct {
}
