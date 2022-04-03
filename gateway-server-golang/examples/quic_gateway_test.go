package main

import (
	gateway "github.com/vanga/quic-gy"
	"github.com/vanga/quic-gy/cmd"
	"testing"
)

func TestUserStory(t *testing.T) {
	cfg := gateway.ParseConfig("/Users/chenhui/github/quic-gateway/gateway-server-golang/cmd/run/gateway.conf")
	si := cmd.NewServer(cfg)
	si.Init()
	si.Start()

	select {}
}
