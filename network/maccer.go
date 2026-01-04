package network

import (
	"fmt"
	"net"
	"net/netip"

	"github.com/mdlayher/arp"
)

func Maccer(ipStr string) {
	iface, err := net.InterfaceByName("eth0")
	if err != nil {
		fmt.Println("Interface introuvable:", err)
		return
	}

	c, err := arp.Dial(iface)
	if err != nil {
		fmt.Println("ARP dial error:", err)
		return
	}
	defer c.Close()

	ip, err := netip.ParseAddr(ipStr)
	if err != nil {
		fmt.Println("IP invalide:", err)
		return
	}

	hw, err := c.Resolve(ip)
	if err != nil {
		fmt.Println("Non résolu")
		return
	}

	fmt.Println("MAC:", hw)
}
