// Print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	occurFiles := make(map[string]map[string]int)
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
				occurFiles[line] = make(map[string]int)
			}
			occurFiles[line][filename] += 1
		}
	}
	for line, occurs := range occurFiles {
		if len(occurs) > 1 {
			fmt.Printf("%s\n", line)
			for filename, n := range occurs {
				fmt.Printf("    %s: %d\n", filename, n)
			}
		}
	}
}
