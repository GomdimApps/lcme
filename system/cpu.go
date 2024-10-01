package system

import (
	"bufio"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type CPUInfo struct {
	NumCores int
	Usage    float64
}

func GetCPUInfo() CPUInfo {
	numCores := runtime.NumCPU()
	usage := getCPUUsage()
	return CPUInfo{NumCores: numCores, Usage: usage}
}

func getCPUUsage() float64 {
	idle0, total0 := getCPUTimes()
	time.Sleep(1 * time.Second)
	idle1, total1 := getCPUTimes()
	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	return 100 * (1 - idleTicks/totalTicks)
}

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
