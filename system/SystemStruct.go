package system

type ServerInfo struct {
	Distribution DistroInfo
	RAM          RAMInfo
	Disk         DiskInfo
	CPU          CPUInfo
	Network      NetworkInfo
	Hardware     HardwareInfo
}
