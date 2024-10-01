package system

import (
	"syscall"
)

type DiskInfo struct {
	Total     uint64
	Used      uint64
	Available uint64
}

func GetDiskInfo(path string) DiskInfo {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return DiskInfo{}
	}
	blockSize := stat.Bsize
	total := (stat.Blocks * uint64(blockSize)) / (1024 * 1024)
	free := (stat.Bfree * uint64(blockSize)) / (1024 * 1024)
	used := total - free
	return DiskInfo{Total: total, Used: used, Available: free}
}
