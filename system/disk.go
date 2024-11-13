package system

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

// DiskInfo is a structure that contains information about the server's disk space.
// It includes the total, used and available space on the system.
// This structure is used by the GetDiskInfo function, which is called in the GetInfoServer function
// function to collect information about the server's disk.
type DiskInfo struct {
	Total     uint64
	Used      uint64
	Available uint64
}

// GetDiskInfo returns information about the disk space on the specified path.
// This function is called inside the GetInfoServer function to get information about the server's disk.
// It uses syscall.Statfs to access file system statistics and calculate space values.
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

// GetFolderSize returns the size of the specified folder in bytes.
func GetFolderSize(path string) (uint64, error) {
	var size uint64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			if os.IsPermission(err) {
				return fmt.Errorf("permission denied: %v", err)
			}
			return err
		}
		if !info.IsDir() {
			size += uint64(info.Size())
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return size / 1024, nil
}
