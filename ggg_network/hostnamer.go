package ggg_network

import (
	"net"
)

func HostnameDiscover(ip string) []string {

	names, _ := net.LookupAddr(ip)
	return names
	
}