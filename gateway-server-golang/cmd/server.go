package cmd

import (
	gateway "github.com/vanga/quic-gy"
	"time"
)

//Server behavior is same as cmd functions
type Server interface {
	//
	Init(cfg *gateway.Config) *Server
	Start() *Server
	Stop(gracefully bool) error
	// Reload graceful reload server, return new Server
	Reload(newCfg *gateway.Config) *Server
	//Upgrade return error & version
	Upgrade() (error, string)
	//Modules list add register modules in Server
	Modules() []*gateway.Module
}

type ServerInstance struct {
	ServerID  string
	cfg       *gateway.Config
	StartTime time.Duration
	pipeline  gateway.Pipeline
}

func (s ServerInstance) Init(cfg *gateway.Config) *Server {
	//TODO implement me
	panic("implement me")
}

func (s ServerInstance) Start() *Server {
	//TODO implement me
	panic("implement me")
}

func (s ServerInstance) Stop(gracefully bool) error {
	//TODO implement me
	panic("implement me")
}

func (s ServerInstance) Reload(newCfg *gateway.Config) *Server {
	//TODO implement me
	panic("implement me")
}

func (s ServerInstance) Upgrade() (error, string) {
	//TODO implement me
	panic("implement me")
}

func (s ServerInstance) Modules() []*gateway.Module {
	//TODO implement me
	panic("implement me")
}
