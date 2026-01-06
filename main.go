package main

import (
    "fmt"
	"log"
	"io"
	"os"
	"flag"
	"strings"
	"sync"
	"bytes"
    "encoding/csv"
	"gopkg.in/yaml.v3"
	"github.com/0adri3n/3g-scan/ggg_network"
)

type Port struct {
	Number string      `yaml:"port"`
	Server string      `yaml:"server,omitempty"`
}

type Machine struct {
	Ip         string   `yaml:"ip"`
	Status     string   `yaml:"status"`
	Hostnames  []string `yaml:"hostnames,omitempty,flow"`
	MacAddress string   `yaml:"macaddress"`
	Vendor     string   `yaml:"vendor"`
	Ports      []Port   `yaml:"ports,omitempty"`
}

func RoutineMaster(ip string, pScan bool, resultsCsvPtr *[][]string, resultsYamlPtr *[]Machine, csvPath string, yamlPath string) {

	up := ggg_network.Pinger(ip)

	var hostnamesStr, macStr, vendorStr, portsStr string

	if up {
		hostnames := ggg_network.HostnameDiscover(ip)
		hostnamesStr = strings.Join(hostnames, "\n")
		m, v := ggg_network.Maccer(ip)
		macStr = m
		vendorStr = v
		if pScan {
			mapped_ports := ggg_network.PortScanner(ip)
			var portBuffer bytes.Buffer
			for key, value := range mapped_ports {
				if value != "" {
					str2write := fmt.Sprintf("%v %v\n", key, value)
					portBuffer.WriteString(str2write)
				} else {
					str2write := fmt.Sprintf("%v\n", key)
					portBuffer.WriteString(str2write)
				}
			}
			portsStr = portBuffer.String()
		}
	}

	var status string
	if up {
		status = "Up"
	} else {
		status = "Down"
	}

	var summaryBuffer bytes.Buffer
	summaryBuffer.WriteString(fmt.Sprintf("\n\nScanning %v\n-----------------------------\n", ip))
	summaryBuffer.WriteString(fmt.Sprintf("Status: %v\n", status))
	summaryBuffer.WriteString(fmt.Sprintf("Hostname(s): %v\n", hostnamesStr))
	summaryBuffer.WriteString(fmt.Sprintf("MAC Address: %v\n", macStr))
	summaryBuffer.WriteString(fmt.Sprintf("Vendor: %v\n", vendorStr))
	summaryBuffer.WriteString(fmt.Sprintf("Ports: \n%v", portsStr))
	log.Println(summaryBuffer.String())

	if csvPath != "" {
		*resultsCsvPtr = append(*resultsCsvPtr, []string{ip, status, hostnamesStr, macStr, vendorStr, portsStr})
	}

	if yamlPath != "" {

		var hostnamesArray []string
		if hostnamesStr != "" {
			hostnamesArray = strings.Split(strings.TrimSpace(hostnamesStr), "\n")
		}

		var portsArray []Port
		if portsStr != "" {
			for _, line := range strings.Split(strings.TrimSpace(portsStr), "\n") {
				line = strings.TrimSpace(line)
				if line == "" {
					continue
				}
				parts := strings.SplitN(line, " Server: ", 2)
				portNum := strings.TrimSpace(parts[0])
				server := ""
				if len(parts) > 1 {
					server = strings.TrimSpace(parts[1])
				}
				portsArray = append(portsArray, Port{Number: portNum, Server: server})
			}
		}

		machine := Machine{
			Ip:         ip,
			Status:     status,
			Hostnames:  hostnamesArray,
			MacAddress: macStr,
			Vendor:     vendorStr,
			Ports:      portsArray,
		}
		*resultsYamlPtr = append(*resultsYamlPtr, machine)
	}

}

func main() {

	log.SetPrefix("3g-scan : ")

	rangesPtr := flag.String("ranges", "", "IP ranges to scan (comma separated)")
	pScanPtr := flag.Bool("p_scan", true, "Port scanning functionality (true/false default true)")
	routinePtr := flag.Bool("routine", true, "Define routines a.k.a threads (true/false default true).")
	debugPtr := flag.Bool("debug", true, "Debug ability (true/false default true)")
	csvPtr := flag.String("csv", "", "CSV output path. If not defined, 3g-scan will not write any CSV file.")
	yamlPtr := flag.String("yaml", "", "YAML output path. If not defined, 3g-scan will not write any yaml file.")

	flag.Parse()

	debug := *debugPtr
	routine := *routinePtr
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
	pScan := *pScanPtr
	csvPath := *csvPtr
	yamlPath := *yamlPtr

	fmt.Println("3g-scan config\n-----------------------------")
	fmt.Println("* IP ranges :")
	for _, ip_range := range ip_ranges {
		fmt.Printf("- %v\n", ip_range)
	}
	fmt.Printf("\n* Port scanning :\n- %v\n", pScan)
	fmt.Printf("\n* Routine :\n- %v\n", routine)
	fmt.Printf("\n* Debug :\n- %v\n", debug)


	fmt.Println("\n-----------------------------")
	fmt.Println("3g-scan started")
	fmt.Println("-----------------------------")


	mapped_ranges := make(map[string][]string)

	for _, ip_range := range ip_ranges {
		listed_ips := ggg_network.CidrLister(ip_range)
		mapped_ranges[ip_range] = listed_ips
	}

	var wg sync.WaitGroup

	resultsCsv := [][]string{

	}

	resultsYaml := []Machine{

	}

	resultsCsvPtr := &resultsCsv
	resultsYamlPtr := &resultsYaml

	for _, ip_range := range ip_ranges {
		ips := mapped_ranges[ip_range]

		for _, ip := range ips {
			
			if routine {

				wg.Add(1)
				go func(ip string) {
					defer wg.Done()
					RoutineMaster(ip, pScan, resultsCsvPtr, resultsYamlPtr, csvPath, yamlPath)
				}(ip)
			} else {
				RoutineMaster(ip, pScan, resultsCsvPtr, resultsYamlPtr, csvPath, yamlPath)
			}


		}

	}

	wg.Wait()

	if csvPath != "" {
		filename := csvPath
		f, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		defer f.Close()
		writer := csv.NewWriter(f)
		writer.Comma = ';' 
		trimSpaces := func(s string) string {
			return strings.TrimSpace(s)
		}
		header := []string{"IP", "Status", "Hostname", "MAC", "Vendor", "Ports"}
		err = writer.Write(header)
		if err != nil {
			fmt.Println("Error writing header:", err)
			return
		}

		for _, result := range resultsCsv {
			trimmedRecord := make([]string, len(result))
			for i, field := range result {
				trimmedRecord[i] = trimSpaces(field)
			}
			err = writer.Write(trimmedRecord)
			if err != nil {
				fmt.Println("Error writing result:", err)
				return
			}
		}
		writer.Flush()
		if err := writer.Error(); err != nil {
			fmt.Println("Error flushing/writing CSV:", err)
		}
	}
	if yamlPath != "" {
		out, err := yaml.Marshal(resultsYaml)
		if err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile(yamlPath, out, 0666); err != nil {
			panic(err)
		}
	}

	fmt.Println("\n-----------------------------")
	fmt.Println("3g-scan done.")
	fmt.Println("-----------------------------")

}
