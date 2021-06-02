package flagx

import (
	"github.com/hsiafan/glow/timex/durationx"
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
	cmd.ParseAndExecute([]string{"-update", "-dry=false", "-name", "kite", "-timeout", "1m", "-tags", "1",
		"-tags", "2", "file1", "3s", "4s"})
	assert.NoError(t, err)
	assert.True(t, op.Update)
	assert.False(t, op.Dry)
	assert.Equal(t, "kite", op.Name)
	assert.Equal(t, []int{1, 2}, op.Tags)
	assert.Equal(t, []string{"4s"}, op.Args)
	assert.Equal(t, "file1", op.File)
	assert.Equal(t, 1, op.Age)
	assert.Equal(t, durationx.Minutes(1), op.Timeout)
	assert.Equal(t, durationx.Seconds(3), op.Timeout2)

	//cmd.ShowUsage()
}

type Option struct {
	Update   bool
	Dry      bool   `default:"true"`
	Name     string `description:"the name"`
	Age      int    `name:"age" description:"the age" default:"1"`
	Tags     []int
	Timeout  time.Duration `default:"0"`
	Args     []string      `flag:"false"`
	File     string        `flag:"false" index:"0"`
	Timeout2 time.Duration `flag:"false" index:"1"`
}
