package gateway

import "runtime/debug"

// ImportPath gateway load module's path
const ImportPath = "github.com/vanga/quic-gy"

//GoModule import caddy-server's module manager
func GoModule() *debug.Module {
	var mod debug.Module
	return goModule(&mod)
}

func goModule(mod *debug.Module) *debug.Module {
	mod.Version = "unknown"
	bi, ok := debug.ReadBuildInfo()
	if ok {
		mod.Path = bi.Main.Path
		for _, dep := range bi.Deps {
			if dep.Path == ImportPath {
				return dep
			}
		}
		return &bi.Main
	}
	return mod
}
