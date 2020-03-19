package flagx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toFlagName(t *testing.T) {
	assert.Equal(t, "test", toFlagName("Test"))
	assert.Equal(t, "test123", toFlagName("Test123"))
	assert.Equal(t, "test-add", toFlagName("TestAdd"))
}

func TestNewCommandLine(t *testing.T) {
	op := &Option{}
	cmd, err := NewCommand("my", "", op)
	assert.NoError(t, err)
	cmd.Description = "a test cmd"
	err = cmd.Parse([]string{"-update", "-dry=false", "-name", "kite", "f1", "f2"})
	assert.NoError(t, err)
	assert.True(t, op.Update)
	assert.False(t, op.Dry)
	assert.Equal(t, "kite", op.Name)
	assert.Equal(t, []string{"f1", "f2"}, op.Files)
	assert.Equal(t, "f2", op.File)
	assert.Equal(t, 1, op.Age)

	//cmd.ShowUsage()
}

type Option struct {
	Update bool
	Dry    bool
	Name   string   `description:"the name"`
	Age    int      `name:"age" description:"the age" default:"1"`
	Files  []string `args:"true"`
	File   string   `args:"true" index:"1"`
}
