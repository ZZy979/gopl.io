// Fetch prints the content found at each specified URL.
//  1. Use io.Copy() instead of io.ReadAll() to copy the response body to os.Stdout without
//     requiring a buffer large enough to hold the entire stream.
//  2. Add the prefix http:// to each argument URL if it is missing.
//  3. Also print the HTTP status code.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Fetched %s\nStatus: %s\n", url, resp.Status)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
