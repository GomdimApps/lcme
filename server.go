package lcme

import (
	"github.com/GomdimApps/lcme/system"
)

type ServerInfo struct {
	Distribution system.DistroInfo
	RAM          system.RAMInfo
	Disk         system.DiskInfo
	CPU          system.CPUInfo
	Network      system.NetworkInfo
}

func GetInfoServer() ServerInfo {
	return ServerInfo{
		Distribution: system.GetDistroInfo(),
		RAM:          system.GetRAMInfo(),
		Disk:         system.GetDiskInfo("/"),
		CPU:          system.GetCPUInfo(),
		Network:      system.GetNetworkInfo(),
	}
}
