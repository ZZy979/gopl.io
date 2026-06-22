package tempconv

import (
	"fmt"
	"testing"
)

func TestArithmetic(t *testing.T) {
	if BoilingC-FreezingC != Celsius(100) {
		t.Error()
	}

	boilingF := CToF(BoilingC)
	if boilingF-CToF(FreezingC) != Fahrenheit(180) {
		t.Error()
	}

	// x := boilingF - FreezingC // compile error: type mismatch
}

func TestComparison(t *testing.T) {
	var c Celsius
	var f Fahrenheit
	testCases := []bool{
		c == 0,
		f >= 0,
		// c == f, // compile error: type mismatch
		c == Celsius(f),
	}
	for _, tc := range testCases {
		if !tc {
			t.Error()
		}
	}
}

func TestString(t *testing.T) {
	c := FToC(212.0)
	testCases := []struct {
		s, expected string
	}{
		{c.String(), "100°C"},
		{fmt.Sprintf("%v", c), "100°C"},
		{fmt.Sprintf("%s", c), "100°C"},
		{fmt.Sprint(c), "100°C"},
		{fmt.Sprintf("%g", c), "100"},
		{fmt.Sprint(float64(c)), "100"},
	}
	for _, tc := range testCases {
		if tc.s != tc.expected {
			t.Errorf("got %v, want %v", tc.s, tc.expected)
		}
	}
}
