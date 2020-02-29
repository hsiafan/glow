package slicex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveIntAt(t *testing.T) {
	var slice = []int{1, 2, 3}
	err := RemoveIntAt(&slice, 2)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2}, slice)

	slice = []int{1, 2, 3}
	err = RemoveIntAt(&slice, 1)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 3}, slice)

	slice = []int{1, 2, 3}
	err = RemoveIntAt(&slice, 0)
	assert.NoError(t, err)
	assert.Equal(t, []int{2, 3}, slice)

	slice = []int{1, 2, 3}
	err = RemoveIntAt(&slice, 3)
	assert.Error(t, err)
}

func TestRemoveAt(t *testing.T) {
	var slice = []int{1, 2, 3}
	err := RemoveAt(&slice, 2)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2}, slice)

	slice = []int{1, 2, 3}
	err = RemoveAt(&slice, 1)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 3}, slice)

	slice = []int{1, 2, 3}
	err = RemoveAt(&slice, 0)
	assert.NoError(t, err)
	assert.Equal(t, []int{2, 3}, slice)

	slice = []int{1, 2, 3}
	err = RemoveAt(&slice, 3)
	assert.Error(t, err)
}

func TestRemoveIntIf(t *testing.T) {
	var slice = []int{1, 2, 3}
	RemoveIntIf(&slice, func(idx int, v int) bool {
		return v >= 3
	})
	assert.Equal(t, []int{1, 2}, slice)

	slice = []int{1, 2, 3}
	RemoveIntIf(&slice, func(idx int, v int) bool {
		return v > 3
	})
	assert.Equal(t, []int{1, 2, 3}, slice)

	slice = []int{1, 2, 3}
	RemoveIntIf(&slice, func(idx int, v int) bool {
		return true
	})
	assert.Equal(t, []int{}, slice)

	slice = []int{}
	RemoveIntIf(&slice, func(idx int, v int) bool {
		return true
	})
	assert.Equal(t, []int{}, slice)

}

func TestRemoveIf(t *testing.T) {
	var slice = []int{1, 2, 3}
	RemoveIf(&slice, func(idx int, v interface{}) bool {
		return v.(int) >= 3
	})
	assert.Equal(t, []int{1, 2}, slice)

	slice = []int{1, 2, 3}
	RemoveIf(&slice, func(idx int, v interface{}) bool {
		return v.(int) > 3
	})
	assert.Equal(t, []int{1, 2, 3}, slice)

	slice = []int{1, 2, 3}
	RemoveIf(&slice, func(idx int, v interface{}) bool {
		return true
	})
	assert.Equal(t, []int{}, slice)

	slice = []int{}
	RemoveIf(&slice, func(idx int, v interface{}) bool {
		return true
	})
	assert.Equal(t, []int{}, slice)
}

func TestFindInt(t *testing.T) {
	assert.Equal(t, 1, FindInt([]int{1, 2, 3}, 2))
	assert.Equal(t, 1, FindInt([]int{1, 2}, 2))
	assert.Equal(t, -1, FindInt([]int{1}, 2))
	assert.Equal(t, -1, FindInt(nil, 2))
}

func TestFind(t *testing.T) {
	assert.Equal(t, 1, Find([]int{1, 2, 3}, 2))
	assert.Equal(t, 1, Find([]int{1, 2}, 2))
	assert.Equal(t, -1, Find([]int{1}, 2))
	assert.Equal(t, -1, Find(nil, 2))
}
