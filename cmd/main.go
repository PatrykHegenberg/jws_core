package main

import (
	"fmt"

	"github.com/PatrykHegenberg/jws_core/pkg/config"
	opsys "github.com/PatrykHegenberg/jws_core/pkg/os"
)

func main() {
	fmt.Println(opsys.GetOperatingSystem())
	config.GetConfig
}
