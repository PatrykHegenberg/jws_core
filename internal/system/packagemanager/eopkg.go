//go:build linux
// +build linux

package packagemanager

import (
	"os/exec"
	"regexp"
	"strings"
)

type Eopkg struct {
	name string
	osid string
}

func NewEopkg(osid string) *Eopkg {
	result := &Eopkg{
		name: "eopkg",
		osid: osid,
	}
	result.intialiseName()
	return result
}

func (e *Eopkg) intialiseName() {
	result := "eopkg"
	stdout, err := exec.Command(".", "eopkg", "--version").Output()
	if err == nil {
		result = strings.TrimSpace(string(stdout))
	}
	e.name = result
}

func (e *Eopkg) InstallCommand(pkg *Package) string {
	if pkg.SystemPackage == false {
		return pkg.InstallCommand[e.osid]
	}
	return "sudo eopkg it " + pkg.Name
}

func (e *Eopkg) Name() string {
	return e.name
}

func (e *Eopkg) PackageInstalled(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	stdout, err := exec.Command(".", "eopkg", "info", pkg.Name).Output()
	return strings.HasPrefix(string(stdout), "Installed"), err
}

func (e *Eopkg) PackageAvailable(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	stdout, err := exec.Command(".", "eopkg", "info", pkg.Name).Output()
	output := e.removeEscapeSequences(string(stdout))
	installed := strings.Contains(output, "Package found in Solus repository")
	e.getPackageVersion(pkg, output)
	return installed, err
}

func (e *Eopkg) removeEscapeSequences(in string) string {
	escapechars, _ := regexp.Compile(`\x1B(?:[@-Z\\-_]|\[[0-?]*[ -/]*[@-~])`)
	return escapechars.ReplaceAllString(in, "")
}

func (e *Eopkg) getPackageVersion(pkg *Package, output string) {
	versionRegex := regexp.MustCompile(`.*Name.*version:\s+(.*)+, release: (.*)`)
	matches := versionRegex.FindStringSubmatch(output)
	pkg.Version = ""
	noOfMatches := len(matches)
	if noOfMatches > 1 {
		pkg.Version = matches[1]
		if noOfMatches > 2 {
			pkg.Version += " (r" + matches[2] + ")"
		}
	}
}
