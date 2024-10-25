package system

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// Address holds the outgoing, incoming, and both (all) addresses.
type Address struct {
	Out []string
	In  []string
	All []string
}

// PortType holds the TCP and UDP addresses.
type PortType struct {
	TCP Address
	UDP Address
}

// NetworkInfo holds the network information.
type NetworkInfo struct {
	IPv4      []string
	IPv6      []string
	IPv4Ports PortType
	IPv6Ports PortType
}

// GetNetworkInfo retrieves the network information.
func GetNetworkInfo() NetworkInfo {
	ipv4s, ipv6s := getIPAddresses()
	return NetworkInfo{
		IPv4: ipv4s,
		IPv6: ipv6s,
		IPv4Ports: PortType{
			TCP: getPortAddresses("/proc/net/tcp"),
			UDP: getPortAddresses("/proc/net/udp"),
		},
		IPv6Ports: PortType{
			TCP: getPortAddresses("/proc/net/tcp6"),
			UDP: getPortAddresses("/proc/net/udp6"),
		},
	}
}

// getIPAddresses retrieves the IPv4 and IPv6 addresses.
func getIPAddresses() ([]string, []string) {
	var ipv4s, ipv6s []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error getting IP addresses:", err)
		return nil, nil
	}

	for _, addr := range addrs {
		ip := extractIP(addr)
		if ip == nil {
			continue
		}
		if ip.To4() != nil {
			ipv4s = append(ipv4s, ip.String())
		} else {
			ipv6s = append(ipv6s, ip.String())
		}
	}
	return ipv4s, ipv6s
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
func getPortAddresses(procFile string) Address {
	var outAddrs, inAddrs, allAddrs []string

	file, err := os.Open(procFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return Address{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Skip the header line
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 10 {
			continue
		}

		localAddr := parseAddress(fields[1])
		remoteAddr := parseAddress(fields[2])

		outAddrs = append(outAddrs, localAddr)
		inAddrs = append(inAddrs, remoteAddr)
		allAddrs = append(allAddrs, fmt.Sprintf("%s > %s", localAddr, remoteAddr))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return Address{
		Out: outAddrs,
		In:  inAddrs,
		All: allAddrs,
	}
}

// parseAddress parses an address in the form "IP:Port".
func parseAddress(addr string) string {
	parts := strings.Split(addr, ":")
	if len(parts) != 2 {
		return ""
	}
	ip := parseIP(parts[0])
	port := parsePort(parts[1])
	return fmt.Sprintf("%s:%d", ip, port)
}

// parseIP parses a hexadecimal IP address.
func parseIP(hexIP string) string {
	if len(hexIP) == 8 { // IPv4
		ip := make(net.IP, 4)
		for i := 0; i < 4; i++ {
			byteVal, _ := strconv.ParseUint(hexIP[i*2:i*2+2], 16, 8)
			ip[3-i] = byte(byteVal)
		}
		return ip.String()
	} else if len(hexIP) == 32 { // IPv6
		ip := make(net.IP, 16)
		for i := 0; i < 16; i++ {
			byteVal, _ := strconv.ParseUint(hexIP[i*2:i*2+2], 16, 8)
			ip[15-i] = byte(byteVal)
		}
		return ip.String()
	}
	return ""
}

// parsePort parses a hexadecimal port number.
func parsePort(hexPort string) uint16 {
	port, _ := strconv.ParseUint(hexPort, 16, 16)
	return uint16(port)
}
