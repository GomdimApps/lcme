package lcme

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/GomdimApps/lcme/system/threads"

	"github.com/GomdimApps/lcme/system"
	"github.com/GomdimApps/lcme/system/utils"
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

// GetFileInfo returns information about the specified files in the given directory.
// If only one file is specified, it returns a slice with one element.
func GetFileInfo(dir string, files ...string) ([]system.FileInfo, error) {
	var fileInfos []system.FileInfo
	// Check if the directory is accessible
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory not found: %v", err)
	} else if err != nil {
		return nil, fmt.Errorf("unable to access directory: %v", err)
	}
	for _, file := range files {
		filePath := filepath.Join(dir, file)
		info, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file not found: %s", file)
		} else if err != nil {
			return nil, fmt.Errorf("error capturing file information: %v", err)
		}
		extension := strings.TrimPrefix(filepath.Ext(file), ".")
		if extension == "" {
			extension = ""
		}
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("error reading file data: %v", err)
		}
		var buffer bytes.Buffer
		buffer.Write(fileData)
		fileInfos = append(fileInfos, system.FileInfo{
			FileName:          info.Name(),
			FileSize:          info.Size() / 1024,
			FileLastChange:    info.ModTime(),
			FileUserPermisson: info.Mode(),
			FileExtension:     extension,
			FileData:          string(fileData),
			FileDataBuffer:    buffer,
			FileDir:           dir,
		})
	}
	return fileInfos, nil
}

// MonitorNetworkRates continuously calculates and returns the download and upload rates.
func MonitorNetworkRates() chan system.NetworkInfo {
	ratesChan := make(chan system.NetworkInfo)
	go func() {
		for {
			initialStats, err := utils.GetNetworkStats()
			if err != nil {
				fmt.Println("Error getting initial network stats:", err)
				continue
			}

			interfaceName, err := utils.GetActiveInterface(initialStats)
			if err != nil {
				fmt.Println("Error getting active interface:", err)
				continue
			}

			downloadRate, uploadRate, err := utils.CalculateNetworkRates(initialStats, interfaceName)
			if err != nil {
				fmt.Println("Error calculating network rates:", err)
				continue
			}

			ratesChan <- system.NetworkInfo{
				Download: downloadRate,
				Upload:   uploadRate,
			}
		}
	}()
	return ratesChan
}

// ScaleFork accepts a task function and manages its execution using the Engine.
func ScaleFork(task threads.Task) {
	engine := threads.NewEngine()
	engine.Start()
	engine.AddTask(task)
	engine.Stop()
}
