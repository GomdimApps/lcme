package lcme

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type ServerInfo struct {
	TotalMemory    uint64
	UsedMemory     uint64
	FreeMemory     uint64
	CPUUsage       float64
	TotalDiskSpace uint64
	UsedDiskSpace  uint64
	FreeDiskSpace  uint64
	Error          error
}

func getInfoServer() ServerInfo {
	var info ServerInfo

	// Memória RAM
	memStats, err := mem.VirtualMemory()
	if err != nil {
		info.Error = err
		return info
	}
	info.TotalMemory = memStats.Total / 1024 / 1024
	info.UsedMemory = memStats.Used / 1024 / 1024
	info.FreeMemory = memStats.Free / 1024 / 1024

	// CPU
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		info.Error = err
		return info
	}
	info.CPUUsage = cpuPercent[0]

	// Disco
	diskUsage, err := disk.Usage("/")
	if err != nil {
		info.Error = err
		return info
	}
	info.TotalDiskSpace = diskUsage.Total / 1024 / 1024
	info.UsedDiskSpace = diskUsage.Used / 1024 / 1024
	info.FreeDiskSpace = diskUsage.Free / 1024 / 1024

	return info
}
