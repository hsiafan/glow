package reflectx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToMap(t *testing.T) {
	var s = struct {
		Test  string
		Test2 int
		test3 int64
	}{
		Test:  "1",
		Test2: 2,
		test3: 10,
	}
	assert.Equal(t, map[string]interface{}{"Test": "1", "Test2": 2}, ToMap(s))
	assert.Equal(t, map[string]interface{}{"Test": "1", "Test2": 2}, ToMap(&s))
}
