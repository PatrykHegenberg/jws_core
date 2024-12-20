package environment

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

type OS struct {
	ID      string
	Name    string
	Version string
}

func Info() (*OS, error) {
	return platformInfo()
}

func GetOperatingSystem() string {
	return runtime.GOOS
}

func platformInfo() (*OS, error) {
	_, err := os.Stat("/etc/os-release")
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("unable to read system information")
	}

	osRelease, _ := os.ReadFile("/etc/os-release")
	return parseOsRelease(string(osRelease)), nil
}

func parseOsRelease(osRelease string) *OS {
	// Default value
	var result OS
	result.ID = "Unknown"
	result.Name = "Unknown"
	result.Version = "Unknown"

	// Split into lines
	lines := strings.Split(osRelease, "\n")
	// Iterate lines
	for _, line := range lines {
		// Split each line by the equals char
		splitLine := strings.SplitN(line, "=", 2)
		// Check we have
		if len(splitLine) != 2 {
			continue
		}
		switch splitLine[0] {
		case "ID":
			result.ID = strings.ToLower(strings.Trim(splitLine[1], "\""))
		case "NAME":
			result.Name = strings.Trim(splitLine[1], "\"")
		case "VERSION_ID":
			result.Version = strings.Trim(splitLine[1], "\"")
		}
	}
	return &result
}
