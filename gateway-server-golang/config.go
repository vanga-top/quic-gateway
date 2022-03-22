package gateway

import (
	"github.com/vanga/quic-gy/utils"
	"io/ioutil"
	"os"
	"strings"
)

// Config for server instance
type Config struct {
	ServerID string

	OriginContent string
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
	//content for config
	content := string(contentByte)
	//start parse
	config := &Config{OriginContent: content}
	/**
	//解析文档的算法
	1、先锁关键字和方法体
	2、
	*/
	strings.Split(content, " ")
	return config
}

type token struct {
}

type liner struct {
	lineContent  string
	lingNum      int
	isStartToken bool
	isEndToken   bool
}
