package lcme

import (
	"github.com/GomdimApps/lcme/system"
	"github.com/GomdimApps/lcme/utils"
)

type ServerInfo struct {
	Distribution system.DistroInfo
	RAM          system.RAMInfo
	Disk         system.DiskInfo
	CPU          system.CPUInfo
	Network      system.NetworkInfo
	Hardware     system.HardwareInfo
}

func GetInfoServer() ServerInfo {
	return ServerInfo{
		Distribution: system.GetDistroInfo(),
		RAM:          system.GetRAMInfo(),
		Disk:         system.GetDiskInfo("/"),
		CPU:          system.GetCPUInfo(),
		Network:      system.GetNetworkInfo(),
		Hardware:     system.GetHardwareInfo(),
	}
}

func Shell(command string) (string, error) {
	return utils.Cexec(command)
}
