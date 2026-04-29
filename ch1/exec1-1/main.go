package main

import (
	"fmt"
	"os"
	"strings"
)

// Print command-line arguments, including name of command.
func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
