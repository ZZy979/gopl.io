package ch4

import (
	"testing"

	"gopl.io/testutil"
)

func TestSha256(t *testing.T) {
	testutil.RunTest(t, "ch4", "sha256", "sha256_output.txt")
}

func TestRev(t *testing.T) {
	testutil.RunTestWithInput(t, "ch4", "rev", "rev_output.txt", "rev_input.txt")
}

func TestAppend(t *testing.T) {
	testutil.RunTest(t, "ch4", "append", "append_output.txt")
}

func TestNonempty(t *testing.T) {
	testutil.RunTest(t, "ch4", "nonempty", "nonempty_output.txt")
}

func TestDedup(t *testing.T) {
	testutil.RunTestWithInput(t, "ch4", "dedup", "dedup_output.txt", "dedup_input.txt")
}

func TestCharcount(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"charcount_input.txt", "charcount_output.txt"},
		{"employee.ser", "charcount_output_invalid.txt"},
	}
	for _, tc := range testCases {
		c := testutil.TestCase{
			Program:    &testutil.Program{"ch4", "charcount"},
			OutputFile: tc.output, InputFile: tc.input, SortLines: true,
		}
		c.Run(t)
	}
}

func TestEmbed(t *testing.T) {
	testutil.RunTest(t, "ch4", "embed", "embed_output.txt")
}

func TestMovie(t *testing.T) {
	testutil.RunTest(t, "ch4", "movie", "movie_output.txt")
}

func TestAutoescape(t *testing.T) {
	testutil.RunTest(t, "ch4", "autoescape", "autoescape.html")
}
