package system

import (
	"fmt"

	"github.com/GomdimApps/lcme/utils"
)

// HardwareInfo is a structure that contains information about the server's hardware,
// such as kernel version, processor name, uptime and swap information.
// This information can be collected by the GetHardwareInfo function and is part of the information returned by GetInfoServer.
type HardwareInfo struct {
	KernelVersion string
	ProcessorName string
	Uptime        int
	SwapTotal     int
	SwapFree      int
}

// GetHardwareInfo is a function that retrieves information about the system's hardware.
// It uses operating system commands to extract data about the kernel version,
// the processor name, the server uptime and the swap area (total and free).
// This data is then returned in the HardwareInfo structure.
// The function is called within GetInfoServer to collect information about the server's hardware.
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

	// Swap Total
	swapTotalStr, err := utils.Cexec("cat /proc/meminfo | grep 'SwapTotal' | uniq | awk '{print $2}'")
	if err != nil {
		fmt.Printf("Error obtaining server Swap Total: %v\n", err)
	}

	// Swap Free
	swapFreeStr, err := utils.Cexec("cat /proc/meminfo | grep 'SwapFree' | uniq | awk '{print $2}'")
	if err != nil {
		fmt.Printf("Error obtaining server Swap free: %v\n", err)
	}

	// Convert String
	ints, err := utils.PassInts(uptimeStr, swapTotalStr, swapFreeStr)
	if err != nil {
		fmt.Printf("Error converting values to integers: %v\n", err)
	}

	return HardwareInfo{
		KernelVersion: kernelVersion,
		ProcessorName: nameCpu,
		Uptime:        ints[0],
		SwapTotal:     ints[1] / 1024,
		SwapFree:      ints[2] / 1024,
	}
}
