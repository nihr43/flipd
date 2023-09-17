package cmd

import (
	"flag"
	"fmt"

	d "flipd/daemon"
)

var cidr string

func Main(args []string) error {
	flag.StringVar(&cidr, "cidr", "10.0.10.0/24", "CIDR address range to be managed by flipd")
	help := flag.Bool("fork", false, "a bool")

	flag.Parse()

	if *help {
		flag.Usage()
		return nil
	}

	fmt.Println("Address range is", cidr)
	return d.Main(cidr)
}
