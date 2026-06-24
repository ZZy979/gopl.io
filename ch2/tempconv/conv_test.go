package tempconv

import (
	"math"
	"testing"
)

func TestConv(t *testing.T) {
	testCases := []struct {
		c Celsius
		f Fahrenheit
		k Kelvin
	}{
		{0, 32, 273.15},
		{1, 33.8, 274.15},
		{100, 212, 373.15},
		{-273.15, -459.67, 0},
	}
	for _, tc := range testCases {
		if f := CToF(tc.c); math.Abs(float64(f-tc.f)) > 1e-6 {
			t.Errorf("CToF(%v) = %v, want %v", tc.c, f, tc.f)
		}
		if c := FToC(tc.f); math.Abs(float64(c-tc.c)) > 1e-6 {
			t.Errorf("FToC(%v) = %v, want %v", tc.f, c, tc.c)
		}
		if k := CToK(tc.c); math.Abs(float64(k-tc.k)) > 1e-6 {
			t.Errorf("CToK(%v) = %v, want %v", tc.c, k, tc.k)
		}
		if c := KToC(tc.k); math.Abs(float64(c-tc.c)) > 1e-6 {
			t.Errorf("KToC(%v) = %v, want %v", tc.k, c, tc.c)
		}
	}
}
