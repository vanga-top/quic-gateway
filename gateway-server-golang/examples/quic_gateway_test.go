package main

import (
	gateway "github.com/vanga/quic-gy"
	"github.com/vanga/quic-gy/cmd"
	"testing"
)

func TestUserStory(t *testing.T) {
	cfg := gateway.ParseConfig("")
	si := &cmd.ServerInstance{}
	si.Init(cfg)
	si.Start()

	select {}
}
