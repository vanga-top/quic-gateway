package cmd

import gateway "github.com/vanga/quic-gy"

//Server behavior is same as cmd functions
type Server interface {
	New(cfg *gateway.Config) *Server
	Start() *Server
	Stop(gracefully bool) error
	Reload(newCfg *gateway.Config) *Server
	//Upgrade return error & version
	Upgrade() (error, string)
	Modules() []*gateway.Module
}
