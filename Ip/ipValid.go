package Ip

import (
	"fmt"
	"net"
	"os"
)

func IpValid() {
	if len(os.Args) != 2 {
		return
	}

	ip := net.ParseIP(os.Args[1])

	if ip != nil {
		fmt.Printf("%v OK\n", ip)
	} else {
		fmt.Println("Bad Address :(")
	}

}
