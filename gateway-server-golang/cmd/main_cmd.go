package cmd

import (
	gateway "github.com/vanga/quic-gy"
	"github.com/vanga/quic-gy/utils"
	"os"
	"strings"

	// plug in Gateway modules here
	_ "github.com/vanga/quic-gy/modules/standard"
)

func init() {
	mod := gateway.GoModule()
	cleanModVersion := strings.TrimPrefix(mod.Version, "v")
	version := "Gateway/" + cleanModVersion
	utils.DefaultLogger.SetLogLevel(utils.LogLevelDebug)
	utils.DefaultLogger.Debugf("version is: %s", version)
}

func Run() {
	switch len(os.Args) {
	case 0:
		utils.DefaultLogger.Errorf("no args be found in cmd...")
		os.Exit(-1)
	case 1:
		os.Args = append(os.Args, "help")
	}
}
