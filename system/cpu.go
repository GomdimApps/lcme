package system

import (
	"bufio"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// CPUInfo contains information about the server's processor.
type CPUInfo struct {
	NumCores int
	Usage    float64
}

// GetCPUInfo returns information about the server's processor, including the number of cores and the current CPU usage.
func GetCPUInfo() CPUInfo {
	return CPUInfo{
		NumCores: runtime.NumCPU(),
		Usage:    getCPUUsage(),
	}
}

// getCPUUsage calculates the current CPU usage by comparing the CPU times at two points in time, with a 1-second interval.
func getCPUUsage() float64 {
	idle0, total0 := getCPUTimes()
	time.Sleep(1 * time.Second)
	idle1, total1 := getCPUTimes()
	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	return 100 * (1 - idleTicks/totalTicks)
}

// getCPUTimes reads the /proc/stat file to get the CPU usage times.
func getCPUTimes() (idle, total uint64) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return 0, 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cpu ") {
			fields := strings.Fields(line)
			return parseCPUTimes(fields)
		}
	}
	return 0, 0
}

// parseCPUTimes parses the CPU times from the fields of the /proc/stat file.
func parseCPUTimes(fields []string) (idle, total uint64) {
	user, _ := strconv.ParseUint(fields[1], 10, 64)
	nice, _ := strconv.ParseUint(fields[2], 10, 64)
	system, _ := strconv.ParseUint(fields[3], 10, 64)
	idle, _ = strconv.ParseUint(fields[4], 10, 64)
	iowait, _ := strconv.ParseUint(fields[5], 10, 64)
	irq, _ := strconv.ParseUint(fields[6], 10, 64)
	softirq, _ := strconv.ParseUint(fields[7], 10, 64)
	total = user + nice + system + idle + iowait + irq + softirq
	return idle, total
}
