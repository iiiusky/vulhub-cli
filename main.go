package main

import (
	"github.com/iiiusky/vulhub-cli/cmd"
	"github.com/iiiusky/vulhub-cli/utils"
)

var (
	version = "0.0.1-beta0"
)

func main() {
	utils.AppVersion = version
	cmd.Execute()
}
