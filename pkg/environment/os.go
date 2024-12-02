package environment

import "runtime"

func GetOperatingSystem() string {
	return runtime.GOOS
}
