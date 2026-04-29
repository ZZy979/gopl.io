package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Measure the difference in running time between echo2 (concatenate strings in loop)
// and echo3 (use strings.Join).
func main() {
	strs := generateRandomStrings(100000, 100)

	start := time.Now()
	result1 := concatWithJoin(strs)
	elapsed := float64(time.Since(start).Nanoseconds()) / 1e6
	fmt.Printf("concatWithJoin costs %.2f ms\n", elapsed)

	start = time.Now()
	result2 := concatWithLoop(strs)
	elapsed = float64(time.Since(start).Nanoseconds()) / 1e6
	fmt.Printf("concatWithLoop costs %.2f ms\n", elapsed)

	fmt.Printf("result1 == result2: %v\n", result1 == result2)
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func generateRandomStrings(count, length int) []string {
	s := make([]string, count)
	for i := range s {
		s[i] = generateRandomString(length)
	}
	return s
}

func concatWithJoin(strs []string) string {
	return strings.Join(strs, " ")
}

func concatWithLoop(strs []string) string {
	result, sep := "", ""
	for _, str := range strs {
		result += sep + str
		sep = " "
	}
	return result
}
