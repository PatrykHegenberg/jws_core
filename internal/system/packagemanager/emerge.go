//go:build linux
// +build linux

package packagemanager

import (
	"os/exec"
	"regexp"
	"strings"
)

type Emerge struct {
	name string
	osid string
}

func NewEmerge(osid string) *Emerge {
	return &Emerge{
		name: "emerge",
		osid: osid,
	}
}

func (e *Emerge) InstallCommand(pkg *Package) string {
	if pkg.SystemPackage == false {
		return pkg.InstallCommand[e.osid]
	}
	return "sudo emerge " + pkg.Name
}

func (e *Emerge) Name() string {
	return e.name
}

func (e *Emerge) PackageInstalled(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	stdout, err := exec.Command(".", "emerge", "-s", pkg.Name+"$").Output()
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok {
			return false, nil
		}
		return false, err
	}

	regex := `.*\*\s+` + regexp.QuoteMeta(pkg.Name) + `\n(?:\S|\s)+?Latest version installed: (.*)`
	installedRegex := regexp.MustCompile(regex)
	matches := installedRegex.FindStringSubmatch(string(stdout))
	pkg.Version = ""
	noOfMatches := len(matches)
	installed := false
	if noOfMatches > 1 && matches[1] != "[ Not Installed ]" {
		installed = true
		pkg.Version = strings.TrimSpace(matches[1])
	}
	return installed, err
}

func (e *Emerge) PackageAvailable(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	stdout, err := exec.Command(".", "emerge", "-s", pkg.Name+"$").Output()
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok {
			return false, nil
		}
		return false, err
	}

	installedRegex := regexp.MustCompile(`.*\*\s+` + regexp.QuoteMeta(pkg.Name) + `\n(?:\S|\s)+?Latest version available: (.*)`)
	matches := installedRegex.FindStringSubmatch(string(stdout))
	pkg.Version = ""
	noOfMatches := len(matches)
	available := false
	if noOfMatches > 1 {
		available = true
		pkg.Version = strings.TrimSpace(matches[1])
	}
	return available, nil
}
