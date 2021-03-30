package vlog

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestFileAppender_Write(t *testing.T) {
	appender, err := NewFileAppender("test_file.log", nil)
	assert.NoError(t, err)
	defer os.Remove("test_file.log")

	err = appender.Append(AppendEvent{"", Debug, "This is a test\n"})
	assert.Nil(t, err)
}

func TestFileAppender_Write2(t *testing.T) {
	appender, err := NewFileAppender("multi/path/test_file.log", nil)
	defer os.RemoveAll("multi/")
	assert.NoError(t, err)

	err = appender.Append(AppendEvent{"", Debug, "This is a test\n"})
	assert.Nil(t, err)
}

func TestGetLogSuffixes(t *testing.T) {
	defer os.RemoveAll("multi/")

	f1, _ := openFile("multi/path/test_file.log")
	f1.Close()
	f4, _ := openFile("multi/path/test_file.log.gz")
	f4.Close()
	f2, _ := openFile("multi/path/test_file.201456.log")
	f2.Close()
	f3, _ := openFile("multi/path/test_file.201457.log.gz")
	f3.Close()

	suffixes := getLogSuffixed("multi/path/test_file.log")
	assert.Equal(t, []string{"201456", "201457"}, suffixes)
}

func TestLogRotate(t *testing.T) {
	defer os.RemoveAll("logs/")
	appender, err := NewFileAppender("logs/test_file.log", nil)
	appender.Append(AppendEvent{"", Debug, "first log\n"})
	assert.NoError(t, err)
	appender.rotateFile("logs/test_file.1234.log")

	_, err = os.Stat("logs/test_file.1234.log")
	assert.NoError(t, err)

	_, err = os.Stat("logs/test_file.log")
	assert.NoError(t, err)

}

func TestTimeRotater(t *testing.T) {
	r := NewHourlyRotater("2006-01-02-15")
	ts, _ := time.Parse("2006-01-02 15:04:05", "2017-05-06 11:12:13")
	r.setInitStatus(ts, 0, []string{})

	b, s := r.Check(ts, 100, 1)
	assert.False(t, b)

	b, s = r.Check(ts.Add(time.Minute), 100, 1)
	assert.False(t, b)

	b, s = r.Check(ts.Add(time.Minute*47), 100, 1)
	assert.False(t, b)

	b, s = r.Check(ts.Add(time.Minute*48), 100, 1)
	assert.True(t, b)
	assert.Equal(t, "2017-05-06-11", s)

	b, s = r.Check(ts.Add(time.Hour), 100, 1)
	assert.False(t, b)
}

func TestSizeRotater(t *testing.T) {
	rotater := NewSizeRotater(1024*1024, 5)
	rotater.setInitStatus(time.Now(), 1024*1023, []string{"xxxx", "123", "0014", "012"})

	b, s := rotater.Check(time.Now(), 1023, 1)
	assert.False(t, b)
	b, s = rotater.Check(time.Now(), 1, 1)
	assert.True(t, b)
	assert.Equal(t, "00124", s)
}
