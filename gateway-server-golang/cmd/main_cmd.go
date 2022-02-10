package cmd

import (
	"fmt"
	gateway "github.com/vanga/quic-gy"
	"strings"

	// plug in Gateway modules here
	_ "github.com/vanga/quic-gy/modules/standard"
)

func init() {
	mod := gateway.GoModule()
	cleanModVersion := strings.TrimPrefix(mod.Version, "v")
	version := "Gateway/" + cleanModVersion
	fmt.Println(version)
}

func Run() {

}
