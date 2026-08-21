package ch3

import (
	"fmt"
	"testing"

	"gopl.io/testutil"
)

func TestBasename(t *testing.T) {
	for i := 1; i <= 2; i++ {
		name := fmt.Sprintf("basename%d", i)
		testutil.RunTestWithInput(t, "ch3", name, "basename_output.txt", "basename_input.txt")
	}
}

func TestComma(t *testing.T) {
	args := []string{"1", "12", "123", "1234", "1234567890", "1234.5678", "一二三四"}
	testutil.RunTest(t, "ch3", "comma", "comma_output.txt", args...)
}

func TestPrintints(t *testing.T) {
	testutil.RunTest(t, "ch3", "printints", "printints_output.txt")
}

func TestNetflag(t *testing.T) {
	testutil.RunTest(t, "ch3", "netflag", "netflag_output.txt")
}
