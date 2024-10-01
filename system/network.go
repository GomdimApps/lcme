package system

import (
	"fmt"
	"net"
)

type NetworkInfo struct {
	IPs []string
}

func GetNetworkInfo() NetworkInfo {
	var ips []string
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting interfaces:", err)
		return NetworkInfo{IPs: ips}
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
				ips = append(ips, ip.String())
			}
		}
	}
	return NetworkInfo{IPs: ips}
}
