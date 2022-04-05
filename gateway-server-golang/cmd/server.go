package cmd

import (
	gateway "github.com/vanga/quic-gy"
	"github.com/vanga/quic-gy/modules/mhttp"
	"time"
)

//Server behavior is same as cmd functions
type Server interface {
	Init() Server
	Start() Server
	Stop(gracefully bool) error
	// Reload graceful reload server, return new Server
	Reload(newCfg *gateway.Config) Server
	//Upgrade return error & version
	Upgrade() (error, string)
	//Modules list add register modules in Server
	Modules() []gateway.Module
	//ServerCMDListener when server cmd event called
	ServerCMDListener(event ServerCMDEvent)
}

type ServerCMDEvent interface {
}

type ServerInstance struct {
	ServerID  string
	cfg       *gateway.Config
	StartTime time.Duration
	pipeline  gateway.Pipeline
}

//New ServerInstance
func NewServer(cfg *gateway.Config) *ServerInstance {
	pi := &gateway.PipelineInstance{}
	return &ServerInstance{pipeline: pi,
		cfg: cfg}
}

func (s *ServerInstance) Init() Server {
	//load all modules
	//build pipeline
	//这里看看能不能解决，不使用不load的模式
	mhModule := mhttp.MHttpModule{}
	s.pipeline.Add(mhModule)
	//TODO implement me
	return s
}

func (s *ServerInstance) Start() Server {
	s.Modules()

	return s
}

func (s *ServerInstance) Stop(gracefully bool) error {
	//TODO implement me
	panic("implement me")
}

func (s *ServerInstance) Reload(newCfg *gateway.Config) Server {
	//TODO implement me
	panic("implement me")
}

func (s *ServerInstance) Upgrade() (error, string) {
	//TODO implement me
	panic("implement me")
}

func (s *ServerInstance) Modules() []gateway.Module {
	return s.pipeline.Chain()
}

func (s *ServerInstance) ServerCMDListener(event ServerCMDEvent) {

}
