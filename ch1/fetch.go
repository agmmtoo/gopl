// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, rawurl := range os.Args[1:] {
		
		var url string

		if strings.HasPrefix(rawurl, "http://") {
			url = rawurl
		} else {
			url = "http://" + rawurl
		}
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%v %s", resp.Status, b)
	}
}
