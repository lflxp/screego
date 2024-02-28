package main

import (
	"github.com/lflxp/screego/cmd"
	pmode "github.com/lflxp/screego/config/mode"
)

var (
	version    = "unknown"
	commitHash = "unknown"
	mode       = pmode.Dev
)

func main() {
	pmode.Set(mode)
	cmd.Run(version, commitHash)
}
