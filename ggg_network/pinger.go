package ggg_network

import (
    "log"
    "time"
    "github.com/go-ping/ping"
)

func Pinger(ip string) bool {
    pinger, err := ping.NewPinger(ip)
    if err != nil {
        panic(err)
    }

    pinger.Count = 3
    pinger.Timeout = time.Second * 3
    pinger.SetPrivileged(true)

    err = pinger.Run()
    if err != nil {
        log.Printf("%v not responding. Skipping...", ip)
        return false
    }

    stats := pinger.Statistics()

    log.Println("Packets sent:", stats.PacketsSent)
    log.Println("Packets received:", stats.PacketsRecv)
    log.Println("Avg RTT:", stats.AvgRtt)

    return true

}