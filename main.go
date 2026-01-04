package main

import (
    "fmt"
	"github.com/0adri3n/3g-scan/network"
)

func main() {
    fmt.Println("Let's scan vicky !")
	ip := "192.168.0.19"
	network.Pinger(ip)
	network.Maccer(ip)
}