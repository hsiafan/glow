package iox

import (
	"io"
	"testing"
)

func TestCloseIfIsCloser(t *testing.T) {
	CloseIfIsCloser(nil)
	var closer io.Closer = nil
	CloseIfIsCloser(closer)
	var closer1 *testClose = nil
	CloseIfIsCloser(closer1)
	var closer2 testClose
	CloseIfIsCloser(&closer2)
}

type testClose int

func (t *testClose) Close() error {
	return nil
}
