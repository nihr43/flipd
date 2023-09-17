package daemon

import (
	"fmt"
)

func Main(cidr string) error {
	fmt.Println("daemon got cidr", cidr)

	err := createbr()
	if err != nil {
		return err
	} else {
		return nil
	}
}
