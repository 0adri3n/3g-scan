package ggg_network

import (
	"log"
	"os/exec"
	"os"
	"regexp"
    "path/filepath"
	"runtime"
	"strings"
	"sync"
	"bufio"
	"bytes"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

var ouiDB map[string]string
var ouiOnce sync.Once

func getOUIDatabase() map[string]string {
	ouiOnce.Do(func() {
		ouiDB = make(map[string]string)
		w_path, err := os.Getwd()
		if err != nil {
			log.Println(err)
			return
		}
		path := filepath.Join(w_path, "/ggg_network/resources/ieee-oui.txt")
		dat, err := os.ReadFile(path)
		if err != nil {
			log.Println("Error reading OUI database:", err)
			return
		}

		scanner := bufio.NewScanner(bytes.NewReader(dat))
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 || strings.HasPrefix(line, "#") {
				continue
			}
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				prefix := strings.ToUpper(parts[0])
				vendor := strings.Join(parts[1:], " ")
				ouiDB[prefix] = vendor
			}
		}

		if err := scanner.Err(); err != nil {
			log.Println("Error scanning OUI database:", err)
		}
	})
	return ouiDB
}

func MacVendor(mac string) string {
	db := getOUIDatabase()
	prefix := strings.ToUpper(strings.ReplaceAll(mac[:8], ":", ""))
	prefix = strings.ToUpper(strings.ReplaceAll(prefix, "-", ""))
	if v, ok := db[prefix]; ok {
		return v
	}
	return "Unknown"
}


func Maccer(ip string) (string, string) {
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

			return m[2], MacVendor(m[2])
		}
	}

	return "MAC not found", "Vendor not found"

}