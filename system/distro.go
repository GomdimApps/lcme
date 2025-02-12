package system

import (
	"fmt"
	"strings"

	"github.com/GomdimApps/lcme/system/utils"
)

// DistroInfo contains information about the operating system distribution.
type DistroInfo struct {
	PrettyName      string
	Name            string
	VersionID       string
	Version         string
	VersionCodeName string
	ID              string
	HomeURL         string
	SupportURL      string
	BugReportURL    string
}

// GetDistroInfo retrieves information about the operating system distribution.
func GetDistroInfo() (DistroInfo, error) {
	output, err := utils.Cexec("cat /etc/os-release")
	if err != nil {
		fmt.Printf("Error retrieving distro info: %v\n", err)
		return unknownDistroInfo(), err
	}

	return parseDistroInfo(output), nil
}

// unknownDistroInfo returns a DistroInfo struct with "Unknown" values.
func unknownDistroInfo() DistroInfo {
	return DistroInfo{
		PrettyName:      "Unknown",
		Name:            "Unknown",
		VersionID:       "Unknown",
		Version:         "Unknown",
		VersionCodeName: "Unknown",
		ID:              "Unknown",
		HomeURL:         "Unknown",
		SupportURL:      "Unknown",
		BugReportURL:    "Unknown",
	}
}

// parseDistroInfo parses the output of /etc/os-release into a DistroInfo struct.
func parseDistroInfo(output string) DistroInfo {
	lines := strings.Split(output, "\n")
	distroInfo := DistroInfo{}

	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.Trim(parts[1], `"`)

		switch key {
		case "PRETTY_NAME":
			distroInfo.PrettyName = value
		case "NAME":
			distroInfo.Name = value
		case "VERSION_ID":
			distroInfo.VersionID = value
		case "VERSION":
			distroInfo.Version = value
		case "VERSION_CODENAME":
			distroInfo.VersionCodeName = value
		case "ID":
			distroInfo.ID = value
		case "HOME_URL":
			distroInfo.HomeURL = value
		case "SUPPORT_URL":
			distroInfo.SupportURL = value
		case "BUG_REPORT_URL":
			distroInfo.BugReportURL = value
		}
	}

	return distroInfo
}
