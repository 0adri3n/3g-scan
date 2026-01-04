package main

import (
    "fmt"
	"log"
	"io"
	"os"
	"flag"
	"runtime"
	"strings"
	"github.com/0adri3n/3g-scan/ggg_network"
)

func main() {

	log.SetPrefix("3g-scan : ")

	rangesPtr := flag.String("ranges", "", "IP ranges to scan (comma separated)")
	ifacePtr := flag.String("iface", "", "Network interface to use (e.g. eth0 or Ethernet)")
	debugPtr := flag.Bool("debug", false, "Debug ability (true/false)")

	flag.Parse()

	debug := *debugPtr
	if debug {
		log.SetOutput(os.Stderr)
		log.SetFlags(log.LstdFlags)
	} else {
		log.SetOutput(io.Discard)
		log.SetFlags(0)
	}

	ranges := *rangesPtr
	ip_ranges := strings.Split(ranges, ",")
	iface := *ifacePtr

	fmt.Println("3g-scan config\n-----------------------------")
	fmt.Println("* IP ranges :")
	for _, ip_range := range ip_ranges {
		fmt.Printf("- %v\n", ip_range)
	}
	fmt.Printf("\n* Interface :\n- %v\n", iface)
	fmt.Printf("\n* Debug :\n- %v", debug)


	fmt.Println("\n-----------------------------")
	fmt.Println("3g-scan started")
	fmt.Println("-----------------------------")


	mapped_ranges := make(map[string][]string)

	for _, ip_range := range ip_ranges {
		listed_ips := ggg_network.CidrLister(ip_range)
		mapped_ranges[ip_range] = listed_ips
	}

	for _, ip_range := range ip_ranges {
		ips := mapped_ranges[ip_range]

		for _, ip := range ips {
			
			log.Printf("\n\nScanning %v\n-----------------------------\n", ip)

			up := ggg_network.Pinger(ip)

			if up {
				switch runtime.GOOS {
				case "windows":
					ggg_network.WindowsMaccer(ip)
				case "linux", "darwin":
					ggg_network.LinuxMaccer(ip, iface)
				}
			}


		}

	}

	var exit string
	fmt.Println("\n\nPress any key then enter to exit...")
	fmt.Scan(&exit)


}