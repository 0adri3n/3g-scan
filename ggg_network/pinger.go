package ggg_network

import (
    "os/exec"
    "runtime"
    "strings"
)

func Pinger(ip string) bool {
    var cmd *exec.Cmd

    if runtime.GOOS == "windows" {
        cmd = exec.Command("ping", "-n", "2", "-w", "2000", ip)
    } else {
        cmd = exec.Command("ping", "-c", "2", "-W", "2", ip)
    }

    output, err := cmd.Output()
    outputStr := string(output)

    if err != nil {
        if !isPacketReceived(outputStr, runtime.GOOS) {
            return false
        }
    }

    if !isPacketReceived(outputStr, runtime.GOOS) {
        return false
    }

    return true
}

func isPacketReceived(output string, osType string) bool {
    outputLower := strings.ToLower(output)
    
    if osType == "windows" {

        hasBytes := strings.Contains(outputLower, "octets=") || strings.Contains(outputLower, "bytes=")
        hasMs := strings.Contains(outputLower, "ms")
        hasTTL := strings.Contains(outputLower, "ttl=")
        
        return hasBytes && (hasMs || hasTTL)
    } else {
        return strings.Contains(outputLower, "time=") || strings.Contains(outputLower, "bytes from")
    }
}
