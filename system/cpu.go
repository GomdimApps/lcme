package system

import (
	"bufio"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// CPUInfo is a structure that contains information about the server's processor.
// It includes the number of CPU cores and the current CPU usage (in percent).
// This structure is used by the GetCPUInfo function, which is called in the GetInfoServer function
// function to collect information about the processor.
type CPUInfo struct {
	NumCores int
	Usage    float64
}

// GetCPUInfo returns information about the server's processor, including the number of cores and the current CPU usage.
// It is called within the GetInfoServer function to obtain information about the server's processor.
func GetCPUInfo() CPUInfo {
	numCores := runtime.NumCPU()
	usage := getCPUUsage()
	return CPUInfo{NumCores: numCores, Usage: usage}
}

// getCPUUsage calculates the current CPU usage by comparing the CPU times at two points in time,
// with a 1 second wait between them.
func getCPUUsage() float64 {
	idle0, total0 := getCPUTimes()
	time.Sleep(1 * time.Second)
	idle1, total1 := getCPUTimes()
	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	return 100 * (1 - idleTicks/totalTicks)
}

// getCPUTimes reads the /proc/stat file to get the CPU usage times.
// This method is used internally to calculate CPU usage in getCPUUsage.
func getCPUTimes() (idle, total uint64) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cpu ") {
			fields := strings.Fields(line)
			user, _ := strconv.ParseUint(fields[1], 10, 64)
			nice, _ := strconv.ParseUint(fields[2], 10, 64)
			system, _ := strconv.ParseUint(fields[3], 10, 64)
			idle, _ = strconv.ParseUint(fields[4], 10, 64)
			iowait, _ := strconv.ParseUint(fields[5], 10, 64)
			irq, _ := strconv.ParseUint(fields[6], 10, 64)
			softirq, _ := strconv.ParseUint(fields[7], 10, 64)
			total = user + nice + system + idle + iowait + irq + softirq
			return
		}
	}
	return
}
