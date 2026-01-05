package ggg_network

import (
	"log"
	"net"
	"time"
	"fmt"
)

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

	for _, port := range commonPorts {

		conn, err := net.DialTimeout(
			"tcp",
			fmt.Sprintf("%s:%d", ip, port),
			500*time.Millisecond,
		)
		if err == nil {
			log.Printf("Port %d: open\n", port)
			conn.Close()
		}

	}

}