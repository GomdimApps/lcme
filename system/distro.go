package system

import (
	"fmt"
	"strings"

	"github.com/GomdimApps/lcme/utils"
)

// DistroInfo is a structure that contains information about the operating system distribution.
// It is used to store relevant system data, such as the name, version, support URL, among others.
// This information can be collected using the GetDistroInfo function and is part of the information that GetInfoServer returns about the server.
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

// GetDistroInfo is a function that retrieves information about the operating system distribution.
// It runs the command `cat /etc/os-release` to get details about the distribution and returns this information
// in the DistroInfo structure. This function is called within GetInfoServer to collect the distribution data from the server.
func GetDistroInfo() (DistroInfo, error) {
	output, err := utils.Cexec("cat /etc/os-release")
	if err != nil {
		fmt.Printf("Error retrieving distro info: %v\n", err)
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
		}, err
	}

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

	return distroInfo, nil
}
