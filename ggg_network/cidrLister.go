package ggg_network

import (
	"net/netip"
	"fmt"
)

func CidrLister(ip_range string) []string {
    var err error
    p, err := netip.ParsePrefix(ip_range)
    if err != nil {
        err = fmt.Errorf("invalid cidr: %s, error %v", ip_range, err)
    }
    
    p = p.Masked()
    
    addr := p.Addr()
	var ips []string
    for {
        if !p.Contains(addr) {
            break
        }
        ips = append(ips, addr.String())
        addr = addr.Next()
    }

	return ips
}
