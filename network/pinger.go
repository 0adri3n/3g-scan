package network

import (
    "fmt"
    "time"
    "github.com/go-ping/ping"
)

func Pinger(ip string) {
    pinger, err := ping.NewPinger(ip)
    if err != nil {
        panic(err)
    }

    pinger.Count = 3
    pinger.Timeout = time.Second * 3
    pinger.SetPrivileged(true) // requis pour ICMP natif

    err = pinger.Run()
    if err != nil {
        panic(err)
    }

    stats := pinger.Statistics()
    fmt.Println("IP:", stats.Addr)
    fmt.Println("Packets sent:", stats.PacketsSent)
    fmt.Println("Packets received:", stats.PacketsRecv)
    fmt.Println("Avg RTT:", stats.AvgRtt)

}