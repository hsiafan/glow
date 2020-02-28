package reflectx

import (
	"path"
	"runtime"
	"strings"
)

// CallerInfo is caller info
type CallerInfo struct {
	Package  string // Package Name
	File     string // File Name
	Function string // Function Name
	LineNo   int    // Line number
}

// GetCaller return the caller info, param depth indicate the stack depth to the caller; for direct caller, depth is 1
func GetCaller(depth int) *CallerInfo {
	pc, file, line, _ := runtime.Caller(depth)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return &CallerInfo{
		Package:  packageName,
		File:     fileName,
		Function: funcName,
		LineNo:   line,
	}
}
