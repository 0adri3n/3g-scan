package ggg_network

import (
	"log"
	"os/exec"
	"regexp"
)


func LinuxMaccer(ip string) {

	cmd := exec.Command("arp", "-a")
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return
	}

	re := regexp.MustCompile(`\((\d+\.\d+\.\d+\.\d+)\)\s+at\s+([0-9a-fA-F:]{17})`)
	matches := re.FindAllStringSubmatch(string(out), -1)

	for _, m := range matches {
		if m[1] == ip {
			log.Printf("MAC Address : %v\n", m[2])
			return
		}
	}

	log.Println("MAC not found")
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
