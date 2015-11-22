// Uses loginstats on a specific set of files.
package main

import (
	"compress/gzip"
	"fmt"
	"github.com/hurstdog/tools/minecraft/loginstats"
	"os"
	"strings"
)

func main() {
	for _, f := range os.Args[1:] {
		fh, err := os.Open(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening %s: %q", f, err)
		}
		if strings.HasSuffix(f, ".gz") {
			gz, err := gzip.NewReader(fh)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading %s: %q", f, err)
			}
			err = loginstats.ReadLog(gz)
		} else {
			err = loginstats.ReadLog(fh)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading %s: %q", f, err)
		}
	}
	stats := loginstats.GetAllStats()
	for k, v := range stats {
		fmt.Printf("Username: %s\n", k)
		fmt.Printf("Login Count: %d\n", v.LoginCount)
		fmt.Printf("Total Play Time: %d minutes\n", v.TotalPlayTime)
	}
}
