package main

import (
    "fmt"
	"log"
	"io"
	"os"
	"flag"
	"strings"
	"github.com/0adri3n/3g-scan/ggg_network"
)

func main() {

	log.SetPrefix("3g-scan : ")

	rangesPtr := flag.String("ranges", "", "IP ranges to scan (comma separated)")
	ifacePtr := flag.String("iface", "", "Network interface to use (e.g. eth0 or Ethernet)")
	pScanPtr := flag.Bool("p_scan", true, "Port scanning functionality (true/false)")
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
	if ranges == "" {
		fmt.Println("Please type in IP ranges to start, using the -ranges argument.")
		return
	}
	ip_ranges := strings.Split(ranges, ",")
	iface := *ifacePtr
	pScan := *pScanPtr

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
				ggg_network.HostnameDiscover(ip)
				ggg_network.Maccer(ip)
				if pScan {
					ggg_network.PortScanner(ip)
				}
			}

		}

	}


	fmt.Println("\n-----------------------------")
	fmt.Println("3g-scan done.")
	fmt.Println("-----------------------------")

	var exit string
	fmt.Println("\n\nPress any key then enter to exit...")
	fmt.Scan(&exit)


}
