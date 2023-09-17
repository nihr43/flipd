package main

import (
	"fmt"
	"os"

	"flipd/cmd"
)

func main() {
	if e := cmd.Main(os.Args); e != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", e)
		os.Exit(1)
	}
}
