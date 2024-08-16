package Ip

import (
	"fmt"
	"net"
	"os"
)

func IpInfo() {
	if len(os.Args) != 2 {
		fmt.Println("Provide the IP Address plz")
		os.Exit(1)
	}

	ip := net.ParseIP(os.Args[1])
	if ip == nil {
		fmt.Println("Unable to parse the IP Address")
		fmt.Println("Ip should be formated using IPv4 or IPv6 notation")
		os.Exit(1)
	}

	fmt.Println()
	fmt.Printf("IP: %s\n", ip)
	fmt.Printf("Default Mask: %s\n", net.IP(ip.DefaultMask()))
	fmt.Printf("Loopback %t\n", ip.IsLoopback())
	fmt.Println("Unicast: ")
	fmt.Printf("Global %t \n", ip.IsGlobalUnicast())
	fmt.Printf("Link %t \n", ip.IsLinkLocalUnicast())
	fmt.Println("Multicast : ")
	fmt.Printf("Global %t \n", ip.IsMulticast())
	fmt.Printf("Interface %t \n", ip.IsInterfaceLocalMulticast())
	fmt.Printf("Link %t \n", ip.IsLinkLocalMulticast())
	fmt.Println()

}
