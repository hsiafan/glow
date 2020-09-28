package flagx

import (
	"github.com/hsiafan/glow/timex"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_toFlagName(t *testing.T) {
	assert.Equal(t, "test", toFlagName("Test"))
	assert.Equal(t, "test123", toFlagName("Test123"))
	assert.Equal(t, "test-add", toFlagName("TestAdd"))
}

func TestNewCommandLine(t *testing.T) {
	op := &Option{}
	cmd, err := NewCommand("my", "", op, func() error {
		return nil
	})
	assert.NoError(t, err)
	cmd.Description = "a test cmd"
	cmd.ParseAndExecute([]string{"-update", "-dry=false", "-name", "kite", "-timeout", "1m", "file1", "3s"})
	assert.NoError(t, err)
	assert.True(t, op.Update)
	assert.False(t, op.Dry)
	assert.Equal(t, "kite", op.Name)
	assert.Equal(t, []string{"file1", "3s"}, op.Args)
	assert.Equal(t, "file1", op.File)
	assert.Equal(t, 1, op.Age)
	assert.Equal(t, timex.MinutesDuration(1), op.Timeout)
	assert.Equal(t, timex.SecondsDuration(3), op.Timeout2)

	//cmd.ShowUsage()
}

type Option struct {
	Update   bool
	Dry      bool
	Name     string        `description:"the name"`
	Age      int           `name:"age" description:"the age" default:"1"`
	Timeout  time.Duration `default:"0"`
	Args     []string      `args:"true"`
	File     string        `args:"true" index:"0"`
	Timeout2 time.Duration `args:"true" index:"1"`
}
