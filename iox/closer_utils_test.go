package iox

import (
	"io"
	"testing"
)

func TestCloseIfIsCloser(t *testing.T) {
	TryClose(nil)
	var closer io.Closer = nil
	TryClose(closer)
	var closer1 *testClose = nil
	TryClose(closer1)
	var closer2 testClose
	TryClose(&closer2)
}

type testClose int

func (t *testClose) Close() error {
	return nil
}
