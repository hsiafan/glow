// +build linux darwin

package vlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSyslogAppender(t *testing.T) {
	appender, err := NewSyslogAppender("vlog")
	assert.NoError(t, err)
	defer appender.Close()
}

func TestSyslogAppender_Append(t *testing.T) {
	appender, _ := NewSyslogAppender("vlog")
	defer appender.Close()
	appender.Append(AppendEvent{"vlog", Info, "This is a test"})
}
