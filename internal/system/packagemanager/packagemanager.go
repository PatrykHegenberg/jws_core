//go:build linux
// +build linux

package packagemanager

import "os/exec"

var pmcommands = []string{
	"eopkg",
	"apt",
	"dnf",
	"pacman",
	"emerge",
	"zypper",
	"nix-env",
}

func Find(osid string) PackageManager {
	for _, pmname := range pmcommands {
		_, err := exec.LookPath(pmname)
		if err == nil {
			return newPackageManager(pmname, osid)
		}
	}
	return nil
}

func newPackageManager(pmname string, osid string) PackageManager {
	switch pmname {
	case "eopkg":
		return NewEopkg(osid)
	case "apt":
		return NewApt(osid)
	case "dnf":
		return NewDnf(osid)
	case "pacman":
		return NewPacman(osid)
	case "emerge":
		return NewEmerge(osid)
	case "zypper":
		return NewZypper(osid)
	case "nix-env":
		return NewNixpkgs(osid)
	}
	return nil
}
