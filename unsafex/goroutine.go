package unsafex

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// CurrentGoroutineId get current goroutine id by walk stack.
func CurrentGoroutineId() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	stk := strings.TrimPrefix(string(buf[:n]), "goroutine ")

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}
