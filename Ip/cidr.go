package Ip

import (
	"flag"
	"fmt"
	"math"
	"net"
	"os"
)

var (
	cidr string
)

func CidrInit() {
	flag.StringVar(&cidr, "c", "", "The CIDR Address")
}

func Cidr() {
	flag.Parse()

	if cidr == "" {
		fmt.Println("CIDR Address is Missing ")
		os.Exit(1)
	}

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println("Failed to parsing CIDR Address", err)
		os.Exit(1)
	}

	// now the calcualtions
	ones, totalBits := ipnet.Mask.Size()
	size := totalBits - ones
	totalHosts := math.Pow(2, float64(size))
	wildCardIp := _wildcard(net.IP(ipnet.Mask))
	last := _lastIp(ip, net.IPMask(wildCardIp))

	fmt.Println()
	fmt.Printf("CIDR %s \n", cidr)
	fmt.Println("--------------------------------------------")
	fmt.Printf("CIDR Block :  %s \n", cidr)
	fmt.Printf("Network : %s \n", ipnet.IP)
	fmt.Printf("Ip Range : %s - %s \n", ip, last)
	fmt.Printf("Total Hosts : %f \n", totalHosts)
	fmt.Printf("Netmask : %s \n", net.IP(ipnet.Mask))
	fmt.Printf("Wildcard Mask : %s \n", wildCardIp)
	fmt.Println()
}

func _wildcard(mask net.IP) net.IP {
	var ipVal net.IP
	for _, octet := range mask {
		ipVal = append(ipVal, ^octet)
	}
	return ipVal
}

func _lastIp(ip net.IP, mask net.IPMask) net.IP {
	ipIn := ip.To4() // v4
	if ipIn == nil {
		ipIn = ip.To16() // v6
		if ipIn == nil {
			return nil
		}
	}

	var ipVal net.IP

	for i, octet := range ipIn {
		ipVal = append(ipVal, octet|mask[i])
	}

	return ipVal
}
