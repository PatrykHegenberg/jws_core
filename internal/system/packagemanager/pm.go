package packagemanager

type Package struct {
	Name           string
	Version        string
	InstallCommand map[string]string
	SystemPackage  bool
	Library        bool
	Optional       bool
}

type packagemap = map[string][]*Package

type PackageManager interface {
	Name() string
	// Packages() packagemap
	PackageInstalled(pkg *Package) (bool, error)
	PackageAvailable(pkg *Package) (bool, error)
	InstallCommand(pkg *Package) string
}

type Dependency struct {
	Name           string
	PackageName    string
	Installed      bool
	InstallCommand string
	Version        string
	Optional       bool
	External       bool
}

type DependencyList []*Dependency

func (d DependencyList) InstallAllRequiredCommand() string {
	result := ""
	for _, dependency := range d {
		if !dependency.Installed && !dependency.Optional {
			result += "  - " + dependency.Name + ": " + dependency.InstallCommand + "\n"
		}
	}

	return result
}

func (d DependencyList) InstallAllOptionalCommand() string {
	result := ""
	for _, dependency := range d {
		if !dependency.Installed && dependency.Optional {
			result += "  - " + dependency.Name + ": " + dependency.InstallCommand + "\n"
		}
	}

	return result
}
