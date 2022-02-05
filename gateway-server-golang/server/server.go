package server

import "context"

//define server ations
type Server interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
	Restart(ctx context.Context)
	Update(version string, ctx context.Context)
}
