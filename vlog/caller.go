package vlog

import (
	"path"
	"runtime"
	"strings"
)

type caller struct {
	packageName  string
	fileName     string
	functionName string
	line         int
}

func getCaller(depth int) *caller {
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

	return &caller{
		packageName:  packageName,
		fileName:     fileName,
		functionName: funcName,
		line:         line,
	}
}
