package system

import (
	"fmt"
	"net"
)

// NetworkInfo is a structure that contains information about the system's network addresses.
// It stores the IPv4 and IPv6 addresses found on the server's network interfaces.
// This information can be collected by the GetNetworkInfo function and is part of the information returned by GetInfoServer.
type NetworkInfo struct {
	IPv4 []string
	IPv6 []string
}

// GetNetworkInfo is a function that retrieves information about the system's network interfaces.
// It uses the net packet to obtain a list of network interfaces and their respective addresses.
// The function checks each network address, ignoring loopback addresses,
// and classifies them as IPv4 or IPv6. It returns this information in the NetworkInfo structure.
// It is called within GetInfoServer to collect information about the server's network.
func GetNetworkInfo() NetworkInfo {

	// Initializes lists to store IPv4 and IPv6 addresses.
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
