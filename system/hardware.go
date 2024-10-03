package system

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/GomdimApps/lcme/utils"
)

type HardwareInfo struct {
	KernelVersion string
	ProcessorName string
	Uptime        int
}

func GetHardwareInfo() HardwareInfo {
	// Kernel version
	kernelVersion, err := utils.Cexec("cat /proc/version | awk '{print $3}'")
	if err != nil {
		fmt.Printf("Error obtaining kernel version: %v\n", err)
	}

	// CPU Name
	nameCpu, err := utils.Cexec("cat /proc/cpuinfo | grep 'model name' | uniq | awk -F ': ' '{print $2}'")
	if err != nil {
		fmt.Printf("Error obtaining CPU core count: %v\n", err)
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
		ProcessorName: nameCpu,
		Uptime:        uptime,
	}
}
