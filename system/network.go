package system

import (
	"fmt"
	"net"
	"strings"

	"github.com/GomdimApps/lcme/utils"
)

// Address represents the network addresses for incoming, outgoing, and all connections.
type Address struct {
	Out []string
	In  []string
	All []string
}

// PortType represents the listening ports and their associated addresses.
type PortType struct {
	TCP Address
	UDP Address
}

// NetworkInfo represents the system's network information, including IPs and ports.
type NetworkInfo struct {
	IPv4      []string
	IPv6      []string
	IPv4Ports PortType
	IPv6Ports PortType
}

// GetNetworkInfo retrieves information about the system's network interfaces and
// the TCP/UDP ports listening on IPv4 and IPv6.
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
			ip := extractIP(addr)
			if ip != nil && !ip.IsLoopback() {
				if ip.To4() != nil {
					ipv4s = append(ipv4s, ip.String())
				} else if ip.To16() != nil {
					ipv6s = append(ipv6s, ip.String())
				}
			}
		}
	}

	return NetworkInfo{
		IPv4: ipv4s,
		IPv6: ipv6s,
		IPv4Ports: PortType{
			TCP: getPortAddresses("ss -4 | grep 'tcp'", "5", "6"),
			UDP: getPortAddresses("ss -4 | grep 'udp'", "5", "6"),
		},
		IPv6Ports: PortType{
			TCP: getPortAddresses("ss -6 | grep 'tcp'", "5", "6"),
			UDP: getPortAddresses("ss -6 | grep 'udp'", "5", "6"),
		},
	}
}

// extractIP extracts the IP address from a net.Addr.
func extractIP(addr net.Addr) net.IP {
	switch v := addr.(type) {
	case *net.IPNet:
		return v.IP
	case *net.IPAddr:
		return v.IP
	}
	return nil
}

// getPortAddresses retrieves the outgoing, incoming, and both (all) addresses for a specific protocol (TCP/UDP).
func getPortAddresses(command, outColumn, inColumn string) Address {
	outAddrs, _ := getPorts(fmt.Sprintf("%s | awk '{print $%s}'", command, outColumn))
	inAddrs, _ := getPorts(fmt.Sprintf("%s | awk '{print $%s}'", command, inColumn))
	allAddrs, _ := getPorts(fmt.Sprintf("%s | awk '{print $%s \" > \" $%s}'", command, outColumn, inColumn))

	return Address{
		Out: outAddrs,
		In:  inAddrs,
		All: allAddrs,
	}
}

// getPorts is a helper function to execute a command and return a list of ports or addresses.
func getPorts(command string) ([]string, error) {
	output, err := utils.Cexec(command)
	if err != nil {
		fmt.Println("Error getting ports:", err)
		return nil, err
	}

	return strings.Split(strings.TrimSpace(output), "\n"), nil
}
