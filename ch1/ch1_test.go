package ch1_test

import (
	"fmt"
	"testing"

	"gopl.io/testutil"
)

func TestHelloWorld(t *testing.T) {
	testutil.RunTest(t, "ch1", "helloworld", "helloworld_output.txt")
}

func TestEcho(t *testing.T) {
	testCases := [][]string{
		{},
		{"hello,", "world!"},
		{"foo bar", "()", "[]", "{}", "!@#$%^&*\\_"},
	}
	for i, tc := range testCases {
		for j := 1; j <= 3; j++ {
			testutil.RunTest(t, "ch1", fmt.Sprintf("echo%d", j),
				fmt.Sprintf("echo_output%d.txt", i+1), tc...)
		}
	}
}

func TestDup1(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"fruits1.txt", "dup_output1.txt"},
		{"fruits2.txt", "dup_output_empty.txt"},
		{"", "dup_output_empty.txt"},
	}
	for _, tc := range testCases {
		c := testutil.TestCase{
			Program:    &testutil.Program{"ch1", "dup1"},
			OutputFile: tc.output, InputFile: tc.input, SortLines: true,
		}
		c.Run(t)
	}
}

func TestDup2(t *testing.T) {
	testCases := []struct {
		args   []string
		input  string
		output string
	}{
		{[]string{"ch1/testdata/fruits1.txt"}, "", "dup_output1.txt"},
		{[]string{"ch1/testdata/fruits1.txt", "ch1/testdata/fruits2.txt"}, "", "dup_output12.txt"},
		{[]string{"ch1/testdata/fruits1.txt", "ch1/testdata/fruits2.txt", "ch1/testdata/fruits3.txt"},
			"", "dup_output123.txt"},
		{[]string{}, "fruits1.txt", "dup_output1.txt"},
		{[]string{}, "", "dup_output_empty.txt"},
	}
	for _, tc := range testCases {
		c := testutil.TestCase{
			Program:    &testutil.Program{"ch1", "dup2"},
			OutputFile: tc.output, Args: tc.args, InputFile: tc.input, SortLines: true,
		}
		c.Run(t)
	}
}

func TestDup3(t *testing.T) {
	testCases := []struct {
		args   []string
		output string
	}{
		{[]string{"ch1/testdata/fruits1.txt"}, "dup_output1.txt"},
		{[]string{"ch1/testdata/fruits1.txt", "ch1/testdata/fruits2.txt"}, "dup_output12.txt"},
		{[]string{"ch1/testdata/fruits1.txt", "ch1/testdata/fruits2.txt", "ch1/testdata/fruits3.txt"},
			"dup_output123.txt"},
		{[]string{}, "dup_output_empty.txt"},
	}
	for _, tc := range testCases {
		c := testutil.TestCase{
			Program:    &testutil.Program{"ch1", "dup3"},
			OutputFile: tc.output, Args: tc.args, SortLines: true,
		}
		c.Run(t)
	}
}
