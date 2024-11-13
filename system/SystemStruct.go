package system

import (
	"os"
	"time"
)

// ServerInfo is a structure that groups together various pieces of information about the server.
// This structure is used to return all the data collected about the server in a single response.
// It contains substructures that store specific information about the system, such as distribution,
// memory, disk, CPU, network and hardware. This information is collected by the GetInfoServer function.
type ServerInfo struct {
	Distribution DistroInfo
	RAM          RAMInfo
	Disk         DiskInfo
	CPU          CPUInfo
	Network      NetworkInfo
	Hardware     HardwareInfo
}

type FileInfo struct {
	FileName          string
	FileSize          int64
	FileLastChange    time.Time
	FileUserPermisson os.FileMode
	FileExtension     string
}
