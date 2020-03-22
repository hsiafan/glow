package flagx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompositeCommand_Parse(t *testing.T) {
	op := &Option{}

	ccmd := NewCompositeCommand("your", "composite command")
	err := ccmd.AddSubCommand("my", "a test cmd", op, func() error {
		assert.True(t, op.Update)
		return nil
	})
	assert.NoError(t, err)
	ccmd.ParseAndExecute([]string{"my", "-update", "f0", "f1"})
	//err = ccmd.ParseAndExecute([]string{"help"})
}
