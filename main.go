package main

import (
    "fmt"
	"runtime"
	"github.com/0adri3n/3g-scan/network"
)

func main() {
	fmt.Println("3g-scan started")

	fmt.Println("Please select an interface : ")
	interfaceName := network.SelectInterface()

    fmt.Println("Let's scan vicky !")

	ip := "192.168.0.23"
	network.Pinger(ip)

	switch runtime.GOOS {
	case "windows":
		network.WindowsMaccer(ip)
	case "linux", "darwin":
		network.LinuxMaccer(ip, interfaceName)
	}

}