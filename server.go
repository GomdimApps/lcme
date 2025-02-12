package lcme

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/GomdimApps/lcme/system"
	"github.com/GomdimApps/lcme/system/compressfiles"
	"github.com/GomdimApps/lcme/system/threads"
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
	engine := threads.NewEngine(runtime.NumCPU())
	engine.Start()
	engine.AddTask(task)
	engine.Stop()
}

// ZipFiles cria um arquivo ZIP contendo os arquivos especificados na slice files.
// Os caminhos salvos no ZIP serão apenas os nomes base dos arquivos.
func ZipFiles(zipFilename string, files []string) error {
	if !strings.HasSuffix(zipFilename, ".zip") {
		return fmt.Errorf("extensão incorreta para arquivo ZIP: %s", zipFilename)
	}

	// Cria o diretório de saída, se necessário.
	outputDir := filepath.Dir(zipFilename)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	newZipFile, err := os.Create(zipFilename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, file := range files {
		// Aqui não passamos baseFolder para usar somente o nome base.
		if err := compressfiles.AddFileToZip(zipWriter, file, ""); err != nil {
			return err
		}
	}
	return nil
}

// TarGzFiles cria um arquivo TAR.GZ contendo os arquivos especificados na slice files.
// Os caminhos salvos serão apenas os nomes base dos arquivos.
func TarGzFiles(tarGzFilename string, files []string) error {
	if !strings.HasSuffix(tarGzFilename, ".tar.gz") {
		return fmt.Errorf("extensão incorreta para arquivo TAR.GZ: %s", tarGzFilename)
	}

	outputDir := filepath.Dir(tarGzFilename)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	tarFile, err := os.Create(tarGzFilename)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	gzipWriter := gzip.NewWriter(tarFile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	for _, file := range files {
		// Aqui também não passamos baseFolder para manter apenas o nome base.
		if err := compressfiles.AddFileToTar(tarWriter, file, ""); err != nil {
			return err
		}
	}
	return nil
}

// ZipFolder cria um arquivo ZIP contendo todos os arquivos (recursivamente)
// dentro da pasta especificada, preservando a estrutura de diretórios.
func ZipFolder(zipFilename, folder string) error {
	if !strings.HasSuffix(zipFilename, ".zip") {
		return fmt.Errorf("extensão incorreta para arquivo ZIP: %s", zipFilename)
	}

	outputDir := filepath.Dir(zipFilename)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	zipFile, err := os.Create(zipFilename)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Caminha pela pasta e adiciona cada arquivo (ignora diretórios).
	err = filepath.Walk(folder, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() {
			return nil
		}
		// Aqui usamos folder como baseFolder para preservar o caminho relativo.
		return compressfiles.AddFileToZip(zipWriter, path, folder)
	})
	return err
}

// TarGzFolder cria um arquivo TAR.GZ contendo todos os arquivos (recursivamente)
// dentro da pasta especificada, preservando a estrutura de diretórios.
func TarGzFolder(tarGzFilename, folder string) error {
	if !strings.HasSuffix(tarGzFilename, ".tar.gz") {
		return fmt.Errorf("extensão incorreta para arquivo TAR.GZ: %s", tarGzFilename)
	}

	outputDir := filepath.Dir(tarGzFilename)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	tarGzFile, err := os.Create(tarGzFilename)
	if err != nil {
		return err
	}
	defer tarGzFile.Close()

	gzipWriter := gzip.NewWriter(tarGzFile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	err = filepath.Walk(folder, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() {
			return nil
		}

		// Check the file size before adding it to the TAR
		if info.Size() > (1<<33)-1 { // 8GB - 1 byte
			return fmt.Errorf("arquivo muito grande para o formato TAR: %s", path)
		}

		return compressfiles.AddFileToTar(tarWriter, path, folder)
	})
	return err
}
