package testutil

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// TestCase 表示示例程序的测试用例
type TestCase struct {
	*Program

	// 用于比较标准输出的文件，相对于章节testdata目录
	OutputFile string

	// 命令行参数（可选）
	Args []string

	// 标准输入文件（可选），相对于章节testdata目录
	InputFile string

	// 是否对输出行进行排序
	SortLines bool
}

// Run 运行指定的程序，并比较标准输出
func (c *TestCase) Run(t *testing.T) {
	t.Helper()

	var stdin io.Reader
	if c.InputFile != "" {
		inputPath := filepath.Join(c.TestDataDir(), c.InputFile)
		input, err := os.ReadFile(inputPath)
		if err != nil {
			t.Fatalf("Failed to open input file %q: %v", inputPath, err)
		}
		stdin = bytes.NewReader(input)
	}

	outputPath := filepath.Join(c.TestDataDir(), c.OutputFile)
	b, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to open output file %q: %v", outputPath, err)
	}
	expectedOutput := replaceNewlines(b)

	actualOutput, err := c.Program.Run(c.Args, stdin)
	if err != nil {
		t.Fatalf("Failed to run test %q: %v", c, err)
	}

	var isMatch bool
	if c.SortLines {
		isMatch = sortLines(string(expectedOutput)) == sortLines(string(actualOutput))
	} else {
		isMatch = bytes.Equal(expectedOutput, actualOutput)
	}
	if !isMatch {
		t.Errorf("Output does not match\nExpected:\n%s\nActual:\n%s", expectedOutput, actualOutput)
	}
}

func replaceNewlines(b []byte) []byte {
	return bytes.ReplaceAll(b, []byte("\r\n"), []byte("\n"))
}

func sortLines(s string) string {
	lines := strings.Split(s, "\n")
	sort.Strings(lines)
	return strings.Join(lines, "\n")
}

func RunTestWithInput(t *testing.T, chapter, name, outputFile, inputFile string, args ...string) {
	c := TestCase{
		Program:    &Program{chapter, name},
		OutputFile: outputFile,
		Args:       args,
		InputFile:  inputFile,
	}
	c.Run(t)
}

func RunTest(t *testing.T, chapter, name, outputFile string, args ...string) {
	RunTestWithInput(t, chapter, name, outputFile, "", args...)
}
