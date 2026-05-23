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

func TestExec1_2(t *testing.T) {
	testCases := [][]string{
		{},
		{"hello,", "world!"},
		{"foo bar", "()", "[]", "{}", "!@#$%^&*\\_"},
	}
	for i, tc := range testCases {
		testutil.RunTest(t, "ch1", "exec1-2",
			fmt.Sprintf("exec1-2_output%d.txt", i+1), tc...)
	}
}

func TestDup1(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"fruits1.txt", "dup_output1.txt"},
		{"fruits2.txt", "empty.txt"},
		{"", "empty.txt"},
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
		{[]string{}, "", "empty.txt"},
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
		{[]string{}, "empty.txt"},
	}
	for _, tc := range testCases {
		c := testutil.TestCase{
			Program:    &testutil.Program{"ch1", "dup3"},
			OutputFile: tc.output, Args: tc.args, SortLines: true,
		}
		c.Run(t)
	}
}

func TestFetch(t *testing.T) {
	testCases := []struct {
		urls   []string
		output string
	}{
		{[]string{"https://httpbin.org/html"}, "moby.html"},
		{[]string{"https://httpbin.org/xml"}, "sample.xml"},
		{[]string{}, "empty.txt"},
	}
	for _, tc := range testCases {
		testutil.RunTest(t, "ch1", "fetch", tc.output, tc.urls...)
	}
}

func TestExec1_7(t *testing.T) {
	testCases := []struct {
		urls   []string
		output string
	}{
		{[]string{"https://httpbin.org/html"}, "exec1-7_output1.txt"},
		{[]string{"httpbin.org/xml"}, "exec1-7_output2.txt"},
		{[]string{"https://httpbin.org/foobar"}, "exec1-7_output3.txt"},
		{[]string{}, "empty.txt"},
	}
	for _, tc := range testCases {
		testutil.RunTest(t, "ch1", "exec1-7", tc.output, tc.urls...)
	}
}
