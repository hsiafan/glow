package flagx

import (
	"errors"
	"fmt"
	"os"
)

// CompositeCommand
type CompositeCommand struct {
	Name        string     // the command name
	Description string     // the description
	subCommands []*Command // sub commands
	hasHelp     bool       // if add help command
}

// Create new CompositeCommand
func NewCompositeCommand(Name string, description string) *CompositeCommand {
	return &CompositeCommand{
		Name:        Name,
		Description: description,
	}
}

// Add one sub command
func (c *CompositeCommand) AddSubCommand(name string, description string, option interface{},
	handle func() error) error {
	command, err := NewCommand(name, description, option, handle)
	if err != nil {
		return err
	}
	command.parentCmd = c.Name
	c.subCommands = append(c.subCommands, command)
	return nil
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
	if len(arguments) == 1 && (arguments[0] == "help" || arguments[0] == "-h" || arguments[0] == "-help") {
		c.ShowUsage()
		return nil
	}
	for _, sc := range c.subCommands {
		if sc.Name == arguments[0] {
			if err := sc.ParseAndExecute(arguments[1:]); err != nil {
				return err
			}
			return nil
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
	for _, command := range c.subCommands {
		fmt.Println("  ", command.Name)
		fmt.Println("    ", command.Description)
	}
}
