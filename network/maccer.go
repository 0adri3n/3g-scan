package network

import (
	"fmt"
	"net"
	"net/netip"
	"os/exec"
	"regexp"

	"github.com/mdlayher/arp"
)

func LinuxMaccer(ipStr string, interfaceName string) {


	iface, err := net.InterfaceByName(interfaceName)
	if err != nil {
		fmt.Println("Interface not found:", err)
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
		fmt.Println("Invalid IP:", err)
		return
	}

	hw, err := c.Resolve(ip)
	if err != nil {
		fmt.Println("Not resolved")
		return
	}

	fmt.Println("MAC:", hw)
}

func WindowsMaccer(ip string) {
	cmd := exec.Command("arp", "-a")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	re := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+)\s+([a-fA-F0-9\-]{17})`)
	matches := re.FindAllStringSubmatch(string(out), -1)

	fmt.Println("Informations from Maccer")

	for _, m := range matches {
		if m[1] == ip {
			fmt.Printf("MAC Adress : %v", m[2])
			return
		}
	}

	fmt.Println("MAC not found")

}
