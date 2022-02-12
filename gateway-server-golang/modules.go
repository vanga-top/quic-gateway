package gateway

import (
	"fmt"
	"net/http"
	"sync"
)

type ModuleID string
type ModuleVersion string

// HandlerFunc is a convenience type like http.HandlerFunc.
type HandlerFunc func(http.ResponseWriter, *http.Request) error

func (f HandlerFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request) error {
	return f(rw, r)
}

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request) error
}

type Module interface {
	ModuleInfo() ModuleInfo
	Handler
	Next() Module
}

type ModuleInfo struct {
	//ID packageName.moduleName.version
	ID      ModuleID
	Version ModuleVersion
	New     func() Module
}

func RegisterModule(module Module) {
	mod := module.ModuleInfo()
	if mod.ID == "" {
		panic("module ID missing")
	}
	if mod.ID == "gateway" || mod.ID == "admin" {
		panic(fmt.Sprintf("module ID '%s' is reserved", mod.ID))
	}
	if mod.New == nil {
		panic("missing ModuleInfo.New")
	}
	if val := mod.New(); val == nil {
		panic("ModuleInfo.New must return a non-nil module instance")
	}
	modulesMu.Lock()
	defer modulesMu.Unlock()

	if _, ok := modules[string(mod.ID)]; ok {
		panic(fmt.Sprintf("module already registered: %s", mod.ID))
	}
	modules[string(mod.ID)] = module
}

func Modules() []Module {

	return nil
}

var (
	modules   = make(map[string]Module)
	modulesMu sync.RWMutex
)
