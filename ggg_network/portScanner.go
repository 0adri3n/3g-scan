package ggg_network

import (
	"log"
	"net"
	"time"
	"fmt"
	"strconv"
	"net/http"
)

func contains(slice []int, element int) bool {
    for _, v := range slice {
        if v == element {
            return true
        }
    }
    return false
}

func HTTPFingerprint(ip string, port int) {
	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get("http://" + ip + ":" + strconv.Itoa(port))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	log.Println("Server:", resp.Header.Get("Server"))
}


func PortScanner(ip string) {
	
    var commonPorts = []int{
        // Réseau & Accès distant
        21,    // FTP
        22,    // SSH
        23,    // Telnet
        3389,  // RDP (Windows)
        5985,  // WinRM HTTP
        5986,  // WinRM HTTPS
        
        // Web
        80,    // HTTP
        443,   // HTTPS
        8080,  // HTTP alt
        8443,  // HTTPS alt
        8000,  // HTTP alt
        8888,  // HTTP alt
        9000,  // HTTP alt
        
        // Email
        25,    // SMTP
        110,   // POP3
        143,   // IMAP
        587,   // SMTP TLS
        
        // DNS & Network
        53,    // DNS
        161,   // SNMP
        389,   // LDAP
        636,   // LDAPS
        
        // Bases de données
        3306,  // MySQL
        5432,  // PostgreSQL
        5984,  // CouchDB
        6379,  // Redis
        7000,  // Cassandra
        27017, // MongoDB
        50070, // Hadoop
        
        // Services
        445,   // SMB (Windows)
        9100,  // Printer
        9200,  // Elasticsearch
    }

	var webPorts = []int {
		// Web
        80,    // HTTP
        443,   // HTTPS
        8080,  // HTTP alt
        8443,  // HTTPS alt
        8000,  // HTTP alt
        8888,  // HTTP alt
        9000,  // HTTP alt
	}

	for _, port := range commonPorts {

		conn, err := net.DialTimeout(
			"tcp",
			fmt.Sprintf("%s:%d", ip, port),
			250*time.Millisecond,
		)
		if err == nil {
			log.Printf("Port %d: open\n", port)
			conn.Close()
			if contains(webPorts, port) {
				HTTPFingerprint(ip, port)
			}
		}

	}

}