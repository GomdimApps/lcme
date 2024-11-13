package lcme

import (
	"fmt"
	"os"
	"strings"

	"github.com/GomdimApps/lcme/system"
	"github.com/GomdimApps/lcme/utils"
)

// GetInfoServer collects detailed information about the server, including
// system distribution, RAM usage, disk space, CPU information,
// network and hardware information.
func GetInfoServer() system.ServerInfo {
	distroInfo, _ := system.GetDistroInfo()
	return system.ServerInfo{
		Distribution: distroInfo,
		RAM:          system.GetRAMInfo(),
		Disk:         system.GetDiskInfo("/"),
		CPU:          system.GetCPUInfo(),
		Network:      system.GetNetworkInfo(),
		Hardware:     system.GetHardwareInfo(),
	}
}

// Shell executes a command in the terminal and returns the result as a string,
// along with an error if one occurs.
func Shell(command string) (string, error) {
	return utils.Cexec(command)
}

// Log returns a log function that writes messages to a specified .log file.
// If the file does not have the .log extension, it displays an error.
func Log(filePath string) func(string) {
	if !strings.HasSuffix(filePath, ".log") {
		fmt.Println("Error: The file must have a .log extension")
		return func(value string) {
			fmt.Println("Error: The file must have a .log extension")
		}
	}

	return func(value string) {
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return
		}
		defer file.Close()

		if _, err := file.WriteString(value + "\n"); err != nil {
			fmt.Println("Error writing to the file:", err)
		}
	}
}

// GetFolderSize returns the size of the specified folder in bytes.
func GetFolderSize(path string) (uint64, error) {
	size, err := system.GetFolderSize(path)
	if err != nil {
		return 0, fmt.Errorf("error getting folder size: %v", err)
	}
	return size, nil
}
