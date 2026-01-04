package ggg_network

import (
	"log"
	"net"
	"net/netip"
	"os/exec"
	"regexp"

	"github.com/mdlayher/arp"
)

func LinuxMaccer(ipStr string, interfaceName string) {


	iface, err := net.InterfaceByName(interfaceName)
	if err != nil {
		log.Printf("Interface not found : %v\n", err)
		return
	}

	c, err := arp.Dial(iface)
	if err != nil {
		log.Printf("ARP dial error : %v\n", err)
		return
	}
	defer c.Close()

	ip, err := netip.ParseAddr(ipStr)
	if err != nil {
		log.Printf("Invalid IP : %v\n", err)
		return
	}

	hw, err := c.Resolve(ip)
	if err != nil {
		log.Println("Not resolved\n")
		return
	}

	log.Printf("MAC address : %v\n", hw)
}

func WindowsMaccer(ip string) {
	cmd := exec.Command("arp", "-a")
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return
	}

	re := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+)\s+([a-fA-F0-9\-]{17})`)
	matches := re.FindAllStringSubmatch(string(out), -1)

	for _, m := range matches {
		if m[1] == ip {
			log.Printf("MAC Address : %v\n", m[2])
			return
		}
	}

	log.Println("MAC not found\n")

}
