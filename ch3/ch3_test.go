package ch3

import (
	"fmt"
	"testing"

	"gopl.io/testutil"
)

func TestSurface(t *testing.T) {
	testutil.RunTest(t, "ch3", "surface", "surface.svg")
}

func TestExec3_1(t *testing.T) {
	testCases := []struct {
		args   []string
		output string
	}{
		{[]string{"-f", "sinc"}, "sinc.svg"},
		{[]string{"-f", "egg-box", "-z", "2", "-zs", "0.1"}, "egg-box.svg"},
		{[]string{"-f", "moguls", "-zs", "0.1"}, "moguls.svg"},
		{[]string{"-f", "saddle", "-c", "50", "-r", "5", "-z", "6.5", "-zs", "0.1"}, "saddle.svg"},
	}
	for _, tc := range testCases {
		testutil.RunTest(t, "ch3", "exec3-1", tc.output, tc.args...)
	}
}

func TestMandelbrot(t *testing.T) {
	c := testutil.TestCase{
		Program:    &testutil.Program{"ch3", "mandelbrot"},
		OutputFile: "mandelbrot.png", Binary: true,
	}
	c.Run(t)
}

func TestExec3_5(t *testing.T) {
	c := testutil.TestCase{
		Program:    &testutil.Program{"ch3", "exec3-5"},
		OutputFile: "full_color_mandelbrot.png", Binary: true,
	}
	c.Run(t)
}

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
