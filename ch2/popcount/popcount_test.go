package popcount

import "testing"

func TestPopCount(t *testing.T) {
	testCases := []struct {
		x        uint64
		expected int
	}{
		{0, 0},
		{1, 1},
		{7, 3},
		{0x80000000, 1},
		{0x9AD6, 9},
		{123456789, 16},
		{0x5555555555555555, 32},
		{0x1234567890ABCDEF, 32},
		{0xDEADBEEFCAFEBABE, 46},
		{0xFFFFFFFFFFFFFFFF, 64},
	}
	algorithms := []struct {
		name string
		f    func(uint64) int
	}{
		{"PopCount", PopCount},
		{"PopCountByLoop", PopCountByLoop},
		{"PopCountByShifting", PopCountByShifting},
		{"PopCountByClearing", PopCountByClearing},
		{"PopCountByDivideAndConquer", PopCountByDivideAndConquer},
	}
	for _, tc := range testCases {
		for _, alg := range algorithms {
			if n := alg.f(tc.x); n != tc.expected {
				t.Errorf("%s(%d) = %d, want %d", alg.name, tc.x, n, tc.expected)
			}
		}
	}
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByDivideAndConquer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByDivideAndConquer(0x1234567890ABCDEF)
	}
}
