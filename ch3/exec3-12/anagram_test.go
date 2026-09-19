package exec3_12

import "testing"

func TestAnagram(t *testing.T) {
	testCases := []struct {
		s1, s2   string
		expected bool
	}{
		{"", "", false},
		{"", "a", false},
		{"a", "a", false},
		{"a", "b", false},
		{"abc", "cba", true},
		{"abc", "abc", false},
		{"abc", "cbd", false},
		{"abc", "abcd", false},
		{"AAAABBBCCD", "ABCDABCABA", true},
		{"BCAAB", "CBBAC", false},
		{"甲乙丙丁戊己庚辛壬癸", "丙癸丁壬庚戊甲乙辛己", true},
	}
	for _, tc := range testCases {
		if r := Anagram(tc.s1, tc.s2); r != tc.expected {
			t.Errorf("Anagram(%q, %q) = %t, want %t", tc.s1, tc.s2, r, tc.expected)
		}
	}
}
