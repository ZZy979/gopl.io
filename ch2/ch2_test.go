package ch2

import (
	"testing"

	"gopl.io/testutil"
)

func TestBoiling(t *testing.T) {
	testutil.RunTest(t, "ch2", "boiling", "boiling_output.txt")
}

func TestFtoc(t *testing.T) {
	testutil.RunTest(t, "ch2", "ftoc", "ftoc_output.txt")
}

func TestEcho4(t *testing.T) {
	testCases := []struct {
		args   []string
		output string
	}{
		{[]string{"a", "bc", "def"}, "echo4_output1.txt"},
		{[]string{"-s", "/", "a", "bc", "def"}, "echo4_output2.txt"},
		{[]string{"-n", "a", "bc", "def"}, "echo4_output3.txt"},
		{[]string{"-s", "##", "-n", "a", "bc", "def"}, "echo4_output4.txt"},
	}
	for _, tc := range testCases {
		testutil.RunTest(t, "ch2", "echo4", tc.output, tc.args...)
	}
}

func TestCf(t *testing.T) {
	args := []string{"32", "212", "-40", "0", "100"}
	testutil.RunTest(t, "ch2", "cf", "cf_output.txt", args...)
}

func TestExec2_2(t *testing.T) {
	args := []string{"1", "2.5", "123.45"}
	testutil.RunTest(t, "ch2", "exec2-2", "exec2-2_output.txt", args...)
}
