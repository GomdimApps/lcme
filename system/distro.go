package system

import (
	"bufio"
	"os"
	"strings"
)

type DistroInfo struct {
	Name string
}

func GetDistroInfo() DistroInfo {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return DistroInfo{Name: "Distribution not found"}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "NAME=") {
			return DistroInfo{Name: strings.Trim(line[6:], "\"")}
		}
	}
	return DistroInfo{Name: "Distribution not found"}
}
