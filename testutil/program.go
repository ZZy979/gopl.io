package testutil

import (
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
)

// Program 表示书中的示例程序
type Program struct {
	// 章节名称，例如"ch1"
	Chapter string

	// 程序名称，例如"helloworld"
	Name string
}

// String 返回程序的完整名称，例如"gopl.io/ch1/helloworld"
func (p *Program) String() string {
	return fmt.Sprintf("gopl.io/%s/%s", p.Chapter, p.Name)
}

// OutDir 返回可执行文件所在目录
func (p *Program) OutDir() string {
	return filepath.Join(OutDir, p.Chapter)
}

// TestDataDir 返回测试数据目录
func (p *Program) TestDataDir() string {
	return filepath.Join(RootDir, p.Chapter, "testdata")
}

// Run 运行示例程序，返回标准输出和错误信息
func (p *Program) Run(args []string, stdin io.Reader) ([]byte, error) {
	cmd := exec.Command(filepath.Join(p.OutDir(), p.Name), args...)
	cmd.Dir = RootDir
	cmd.Stdin = stdin
	return cmd.CombinedOutput()
}
