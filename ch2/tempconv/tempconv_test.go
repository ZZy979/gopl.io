package tempconv

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	testCases := []struct {
		s, expected string
	}{
		{FreezingC.String(), "0°C"},
		{fmt.Sprintf("%v", BoilingC), "100°C"},
		{fmt.Sprintf("%s", AbsoluteZeroC), "-273.15°C"},
		{fmt.Sprint(Fahrenheit(32)), "32°F"},
		{Fahrenheit(-300).String(), "-300°F"},
		{Kelvin(4000).String(), "4000K"},
	}
	for _, tc := range testCases {
		if tc.s != tc.expected {
			t.Errorf("got %v, want %v", tc.s, tc.expected)
		}
	}
}
