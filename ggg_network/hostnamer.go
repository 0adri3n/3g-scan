package ggg_network

import (
	"log"
	"net"
)

func HostnameDiscover(ip string) []string {

	names, _ := net.LookupAddr(ip)
    log.Println("Hostname(s) :", names)
	return names
	
}