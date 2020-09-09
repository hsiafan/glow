// +build linux darwin

package vsyslog

import (
	"github.com/hsiafan/glow/vlog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSyslogAppender(t *testing.T) {
	appender, err := New("vlog")
	assert.NoError(t, err)
	defer appender.Close()
}

func TestSyslogAppender_Append(t *testing.T) {
	appender, _ := New("vlog")
	defer appender.Close()
	appender.Append(vlog.AppendEvent{"vlog", vlog.Info, "This is a test"})
}
