package system

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// RAMInfo is a structure that contains information about the system's RAM memory.
// It is used to store data about the total, usage and availability of RAM memory.
// This information can be collected by the GetRAMInfo function and is part of the information returned by GetInfoServer.
type RAMInfo struct {
	Total     uint64
	Used      uint64
	Available uint64
}

// GetRAMInfo is a function that retrieves information about the system's RAM memory.
// It reads the `/proc/meminfo` file to obtain data on total and available memory,
// and calculates the memory used. The function returns this data in the RAMInfo structure.
// It is called inside GetInfoServer to collect information about the server's RAM.
func GetRAMInfo() RAMInfo {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return RAMInfo{}
	}
	defer file.Close()

	var totalRAM, availableRAM uint64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if strings.HasPrefix(line, "MemTotal:") {
			totalRAM, _ = strconv.ParseUint(fields[1], 10, 64)
		}
		if strings.HasPrefix(line, "MemAvailable:") {
			availableRAM, _ = strconv.ParseUint(fields[1], 10, 64)
			break
		}
	}

	totalRAMMB := totalRAM / 1024
	availableRAMMB := availableRAM / 1024
	usedRAMMB := totalRAMMB - availableRAMMB
	return RAMInfo{Total: totalRAMMB, Used: usedRAMMB, Available: availableRAMMB}
}
