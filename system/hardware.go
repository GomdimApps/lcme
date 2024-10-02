package system

import (
	"fmt"
	"strconv"

	"github.com/GomdimApps/lcme/utils"
)

type HardwareInfo struct {
	KernelVersion string
	NumCores      string
	CPUMHz        string
	Uptime        int
}

func GetHardwareInfo() HardwareInfo {
	// kernel version
	kernelVersion, err := utils.Cexec("cat /proc/version | awk '{print $3}'")
	if err != nil {
		fmt.Printf("Error getting the kernel version: %v\n", err)
	}

	// CPU cores
	numCores, err := utils.Cexec("cat /proc/cpuinfo | grep 'cpu cores' | uniq | awk -F ': ' '{print $2}'")
	if err != nil {
		fmt.Printf("Error getting the number of CPU cores: %v\n", err)
	}

	// CPU frequency (in MHz)
	cpuMHz, err := utils.Cexec("cat /proc/cpuinfo | grep 'cpu MHz' | uniq | awk -F ': ' '{print $2}'")
	if err != nil {
		fmt.Printf("Error obtaining the CPU frequency: %v\n", err)
	}

	// Server uptime
	uptimeStr, err := utils.Cexec("awk '{print int($1/60)}' /proc/uptime")
	if err != nil {
		fmt.Printf("Error getting the active time from the server: %v\n", err)
	}

	uptime, err := strconv.Atoi(uptimeStr)
	if err != nil {
		fmt.Printf("Error converting uptime to integer:  %v\n", err)
	}

	return HardwareInfo{
		KernelVersion: kernelVersion,
		NumCores:      numCores,
		CPUMHz:        cpuMHz,
		Uptime:        uptime,
	}
}
