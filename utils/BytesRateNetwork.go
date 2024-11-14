package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// GetNetworkStats reads the network statistics from /proc/net/dev and returns a map of interface names to their received and transmitted bytes.
func GetNetworkStats() (map[string][2]int64, error) {
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats := make(map[string][2]int64)
	scanner := bufio.NewScanner(file)
	scanner.Scan() // Skip the first header line
	scanner.Scan() // Skip the second header line

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 10 {
			continue
		}

		iface := strings.TrimSuffix(fields[0], ":")
		rxBytes, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			return nil, err
		}
		txBytes, err := strconv.ParseInt(fields[9], 10, 64)
		if err != nil {
			return nil, err
		}

		stats[iface] = [2]int64{rxBytes, txBytes}
	}

	return stats, nil
}

// GetActiveInterface returns the name of the active network interface.
func GetActiveInterface(stats map[string][2]int64) (string, error) {
	maxBytes := int64(0)
	activeInterface := ""
	for iface, bytes := range stats {
		if bytes[0] > maxBytes {
			maxBytes = bytes[0]
			activeInterface = iface
		}
	}

	if activeInterface == "" {
		return "", fmt.Errorf("no active interface found")
	}
	return activeInterface, nil
}

// CalculateNetworkRates calculates the download and upload rates for the active network interface.
func CalculateNetworkRates() (downloadRate, uploadRate int64, err error) {
	initialStats, err := GetNetworkStats()
	if err != nil {
		return 0, 0, err
	}

	interfaceName, err := GetActiveInterface(initialStats)
	if err != nil {
		return 0, 0, err
	}

	initialBytes := initialStats[interfaceName]

	time.Sleep(1 * time.Second)

	finalStats, err := GetNetworkStats()
	if err != nil {
		return 0, 0, err
	}

	finalBytes := finalStats[interfaceName]

	downloadRate = (finalBytes[0] - initialBytes[0]) / 1024 // Calculate in KBps
	uploadRate = (finalBytes[1] - initialBytes[1]) / 1024   // Calculate in KBps

	return downloadRate, uploadRate, nil
}
