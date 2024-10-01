package system

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type RAMInfo struct {
	Total     uint64
	Used      uint64
	Available uint64
}

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
