package testutil

import (
	"path/filepath"
	"runtime"
)

var (
	RootDir = findProjectRoot()
	OutDir  = filepath.Join(RootDir, "out")
)

// 查找项目根目录
func findProjectRoot() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(file))
}
