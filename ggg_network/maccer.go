package ggg_network

import (
	"log"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)


func MacVendor(mac string, db map[string]string) string {
	prefix := strings.ToUpper(strings.ReplaceAll(mac[:8], ":", ""))
	prefix = strings.ToUpper(strings.ReplaceAll(prefix, "-", ""))
	if v, ok := db[prefix]; ok {
		return v
	}
	return "Unknown"
}


func Maccer(ip string, db map[string]string) (string, string) {
	cmd := exec.Command("arp", "-a")
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return "", ""
	}

	var re *regexp.Regexp

	switch runtime.GOOS {
	case "windows":
		re = regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+)\s+([a-fA-F0-9\-]{17})`)
	case "linux", "darwin":
		re = regexp.MustCompile(`\((\d+\.\d+\.\d+\.\d+)\)\s+at\s+([0-9a-fA-F:]{17})`)
	}

	matches := re.FindAllStringSubmatch(string(out), -1)

	for _, m := range matches {
		if m[1] == ip {

			return m[2], MacVendor(m[2], db)
		}
	}

	return "MAC not found", "Vendor not found"

}