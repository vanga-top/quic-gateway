package gateway

import (
	"fmt"
	"github.com/vanga/quic-gy/utils"
	"io/ioutil"
	"os"
)

// Config for server instance
type Config struct {
	ServerID string
}

// ParseConfig parse gateway.conf to Config
func ParseConfig(path string) *Config {
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		utils.DefaultLogger.Errorf("error in parse config: %s", err)
	}
	contentByte, ioerror := ioutil.ReadAll(f)
	if ioerror != nil {
		utils.DefaultLogger.Errorf("error in io read: %s", ioerror)
	}
	fmt.Println(string(contentByte))
	
	return nil
}
