package daemon

import (
	"github.com/vishvananda/netlink"
)

func createbr() error {
    la := netlink.NewLinkAttrs()
    la.Name = "foo"
    mybridge := &netlink.Bridge{LinkAttrs: la}
    err := netlink.LinkAdd(mybridge)
    if err != nil  {
	return err
    } else {
        return nil
    }
}
