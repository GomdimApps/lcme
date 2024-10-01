package system

import (
	"fmt"
	"net"
)

type NetworkInfo struct {
	IPv4 []string
	IPv6 []string
}

func GetNetworkInfo() NetworkInfo {
	var ipv4s, ipv6s []string

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting interfaces:", err)
		return NetworkInfo{IPv4: ipv4s, IPv6: ipv6s}
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip != nil && !ip.IsLoopback() {
				if ip.To4() != nil {
					ipv4s = append(ipv4s, ip.String())
				} else if ip.To16() != nil {
					ipv6s = append(ipv6s, ip.String())
				}
			}
		}
	}

	return NetworkInfo{IPv4: ipv4s, IPv6: ipv6s}
}
