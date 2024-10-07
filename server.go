package lcme

import (
	"fmt"
	"os"
	"strings"

	"github.com/GomdimApps/lcme/system"
	"github.com/GomdimApps/lcme/utils"
)

func GetInfoServer() system.ServerInfo {
	return system.ServerInfo{
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
