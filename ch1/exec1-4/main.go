package main

import (
	"bufio"
	"fmt"
	"os"
)

// Print the names of all files in which each duplicated line occurs.
func main() {
	occurFiles := make(map[string]map[string]bool)
	for _, filename := range os.Args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			line := input.Text()
			if _, ok := occurFiles[line]; !ok {
				occurFiles[line] = make(map[string]bool)
			}
			occurFiles[line][filename] = true
		}
	}
	for line, occurs := range occurFiles {
		if len(occurs) > 1 {
			fmt.Printf("%s\n", line)
			for filename := range occurs {
				fmt.Printf("\t%s\n", filename)
			}
		}
	}
}
