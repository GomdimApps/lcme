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
		Usage:    calculateCPUUsage(),
	}
}

// calculateCPUUsage calculates the current CPU usage by taking multiple samples over a period of time and averaging them.
func calculateCPUUsage() float64 {
	const sampleCount = 5
	var idleTotal, totalTotal uint64

	for i := 0; i < sampleCount; i++ {
		idle, total := readCPUTimes()
		idleTotal += idle
		totalTotal += total
		time.Sleep(200 * time.Millisecond)
	}

	idleAvg := idleTotal / sampleCount
	totalAvg := totalTotal / sampleCount

	return computeUsage(idleAvg, totalAvg)
}

// readCPUTimes reads the /proc/stat file to get the CPU usage times.
func readCPUTimes() (idle, total uint64) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return 0, 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cpu ") {
			return parseCPUTimes(strings.Fields(line))
		}
	}
	return 0, 0
}

// parseCPUTimes parses the CPU times from the fields of the /proc/stat file.
func parseCPUTimes(fields []string) (idle, total uint64) {
	var times [7]uint64
	for i := 1; i <= 7; i++ {
		times[i-1], _ = strconv.ParseUint(fields[i], 10, 64)
	}
	idle = times[3]
	total = sum(times[:])
	return idle, total
}

// sum calculates the sum of a slice of uint64.
func sum(values []uint64) uint64 {
	var total uint64
	for _, value := range values {
		total += value
	}
	return total
}

// computeUsage calculates the CPU usage percentage.
func computeUsage(idle, total uint64) float64 {
	idleTicks := float64(idle)
	totalTicks := float64(total)
	return 100 * (1 - idleTicks/totalTicks)
}
