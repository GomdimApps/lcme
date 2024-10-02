package system

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/GomdimApps/lcme/utils"
)

type HardwareInfo struct {
	KernelVersion string
	NumCores      string
	CPUMHz        string
	Uptime        int
}

func GetHardwareInfo() HardwareInfo {
	// Kernel version
	kernelVersion, err := utils.Cexec("cat /proc/version | awk '{print $3}'")
	if err != nil {
		fmt.Printf("Error obtaining kernel version: %v\n", err)
	}

	// CPU cores
	numCores, err := utils.Cexec("cat /proc/cpuinfo | grep 'cpu cores' | uniq | awk -F ': ' '{print $2}'")
	if err != nil {
		fmt.Printf("Error obtaining CPU core count: %v\n", err)
	}

	// CPU frequency (in MHz)
	cpuMHz, err := utils.Cexec("cat /proc/cpuinfo | grep 'cpu MHz' | uniq | awk -F ': ' '{print $2}'")
	if err != nil {
		fmt.Printf("Error obtaining CPU frequency: %v\n", err)
	}

	// Server uptime
	uptimeStr, err := utils.Cexec("awk '{print int($1/60)}' /proc/uptime")
	if err != nil {
		fmt.Printf("Error obtaining server uptime: %v\n", err)
	}

	uptimeStr = strings.TrimSpace(uptimeStr)
	uptime, err := strconv.Atoi(uptimeStr)
	if err != nil {
		fmt.Printf("Error converting uptime to integer: %v\n", err)
	}

	return HardwareInfo{
		KernelVersion: kernelVersion,
		NumCores:      numCores,
		CPUMHz:        cpuMHz,
		Uptime:        uptime,
	}
}
