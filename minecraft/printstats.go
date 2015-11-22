// Uses loginstats on a specific set of files.
package main

import (
	"fmt"
	"github.com/hurstdog/tools/minecraft/loginstats"
	"os"
)

func main() {
	for _, f := range os.Args[1:] {
		err := loginstats.ReadLog(f)
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
