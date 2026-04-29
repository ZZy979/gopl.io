package main

import (
	"fmt"
	"os"
)

// Print the index and value of each command-line argument, one per line.
func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%d: %s\n", i, os.Args[i])
	}
}
