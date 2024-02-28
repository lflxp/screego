package main

import (
	"github.com/lflxp/srceego/cmd"
	pmode "github.com/lflxp/srceego/config/mode"
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
