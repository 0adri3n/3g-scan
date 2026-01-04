package network

import (
	"fmt"
	"net"
)

func SelectInterface() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error on interfaces reading:", err)
		return ""
	}

	for index, iface := range interfaces {
		fmt.Printf("ID : %v | Interface informations : %v\n", index, iface)
	}

	var selectedIndex int;

	fmt.Println("Please select interface ID : ")
	fmt.Scan(&selectedIndex)

	selectedIface := interfaces[selectedIndex].Name

	return selectedIface
}