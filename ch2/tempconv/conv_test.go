package tempconv

import (
	"math"
	"testing"
)

func TestConv(t *testing.T) {
	testCases := []struct {
		c Celsius
		f Fahrenheit
	}{
		{0, 32},
		{1, 33.8},
		{100, 212},
		{-273.15, -459.67},
	}
	for _, tc := range testCases {
		if f := CToF(tc.c); math.Abs(float64(f-tc.f)) > 1e-6 {
			t.Errorf("CToF(%v) = %v, want %v", tc.c, f, tc.f)
		}
		if c := FToC(tc.f); math.Abs(float64(c-tc.c)) > 1e-6 {
			t.Errorf("FToC(%v) = %v, want %v", tc.f, c, tc.c)
		}
	}
}
