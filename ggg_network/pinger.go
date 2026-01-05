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

    pinger.Count = 2
    pinger.Timeout = time.Second * 2
    pinger.SetPrivileged(true)

    err = pinger.Run()
    if err != nil {
        log.Printf("%v not responding. Skipping...", ip)
        return false
    }

    stats := pinger.Statistics()

    var status string

    if stats.PacketsRecv == 0 {
        status = "Down"
    } else {
        status = "Up"
    }

    log.Println("Packets sent :", stats.PacketsSent)
    log.Println("Packets received :", stats.PacketsRecv)
    log.Println("Avg RTT :", stats.AvgRtt)
    log.Println("Status :", status)

    return true

}