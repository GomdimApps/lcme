package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// GetActiveInterface returns the name of the active network interface.
func GetActiveInterface() (string, error) {
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Skip the first header line
	scanner.Scan() // Skip the second header line

	maxBytes := int64(0)
	activeInterface := ""

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) < 10 {
			continue
		}

		iface := strings.TrimSuffix(fields[0], ":")
		rxBytes, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			return "", err
		}

		if rxBytes > maxBytes {
			maxBytes = rxBytes
			activeInterface = iface
		}
	}

	if activeInterface == "" {
		return "", fmt.Errorf("no active interface found")
	}
	return activeInterface, nil
}

// GetBytes returns the received and transmitted bytes for a given network interface.
func GetBytes(interfaceName string) (rxBytes, txBytes int64, err error) {
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, interfaceName) {
			fields := strings.Fields(line)
			if len(fields) < 10 {
				return 0, 0, fmt.Errorf("unexpected format in /proc/net/dev")
			}
			rxBytes, err = strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				return 0, 0, err
			}
			txBytes, err = strconv.ParseInt(fields[9], 10, 64)
			if err != nil {
				return 0, 0, err
			}
			return rxBytes, txBytes, nil
		}
	}
	return 0, 0, fmt.Errorf("interface %s not found", interfaceName)
}

// CalculateNetworkRates calculates the download and upload rates for the active network interface.
func CalculateNetworkRates() (downloadRate, uploadRate int64, err error) {
	interfaceName, err := GetActiveInterface()
	if err != nil {
		return 0, 0, err
	}

	rxBytesInitial, txBytesInitial, err := GetBytes(interfaceName)
	if err != nil {
		return 0, 0, err
	}

	time.Sleep(1 * time.Second)

	rxBytesFinal, txBytesFinal, err := GetBytes(interfaceName)
	if err != nil {
		return 0, 0, err
	}

	downloadRate = (rxBytesFinal - rxBytesInitial) / 1024 // Calculate in KBps
	uploadRate = (txBytesFinal - txBytesInitial) / 1024   // Calculate in KBps

	return downloadRate, uploadRate, nil
}
