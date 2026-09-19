package exec3_12

import (
	"slices"
)

// Anagram reports whether two strings are anagrams of each other,
// that is, they contain the same letters in a different order.
func Anagram(s1, s2 string) bool {
	if s1 == s2 || len(s1) != len(s2) {
		return false
	}
	r1, r2 := []rune(s1), []rune(s2)
	slices.Sort(r1)
	slices.Sort(r2)
	return slices.Equal(r1, r2)
}
