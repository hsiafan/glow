package flagx

import (
	"fmt"
	"os"
)

// CompositeCommand
type CompositeCommand struct {
	Name        string     // the command name
	Description string     // the description
	subCommands []*Command // sub commands
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
	handle Handle) error {
	command, err := NewCommand(name, description, option, handle)
	if err != nil {
		return err
	}
	command.parentCmd = c.Name
	c.subCommands = append(c.subCommands, command)
	return nil
}

// Parse commandline passed arguments, and execute command
func (c *CompositeCommand) ParseOsArgsAndExecute() {
	c.ParseAndExecute(os.Args[1:])
}

// Parse arguments, and execute command
func (c *CompositeCommand) ParseAndExecute(arguments []string) {
	if len(arguments) == 0 {
		arguments = []string{"help"}
	}
	if len(arguments) == 1 && (arguments[0] == "help" || arguments[0] == "-h" || arguments[0] == "-help") {
		c.ShowUsage()
		return
	}
	for _, sc := range c.subCommands {
		if sc.Name == arguments[0] {
			sc.ParseAndExecute(arguments[1:])
			return
		}
	}
	fmt.Println("unknown command: " + arguments[0])
	os.Exit(-1)
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
