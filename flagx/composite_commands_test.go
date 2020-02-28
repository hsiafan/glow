package flagx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompositeCommand_Parse(t *testing.T) {
	op := &Option{}
	cmd, err := NewCommand("my", "", op)
	assert.NoError(t, err)
	cmd.Description = "a test cmd"

	ccmd := NewCompositeCommand("your")
	ccmd.AddHelpCommand()
	ccmd.AddSubCommand(cmd, func() error {
		assert.True(t, op.Update)
		return nil
	})
	err = ccmd.ParseAndExecute([]string{"my", "-update", "f0", "f1"})
	//err = ccmd.ParseAndExecute([]string{"help"})
	assert.NoError(t, err)
}
