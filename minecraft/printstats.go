// Uses loginstats on a specific set of files.
package main

import (
	"bufio"
	"fmt"
	"loginstats"
	"os"
)

func main() {
	for _, f := range os.Args[1:] {
		loginstats.ReadLog(f)
	}
}
