package mhttp

import (
	"fmt"
	gateway "github.com/vanga/quic-gy"
	"net/http"
)

func init() {
	fmt.Printf("load mhttp module....")
}

type MHttpModule struct {
}

func (M MHttpModule) ModuleInfo() gateway.ModuleInfo {
	//TODO implement me
	panic("implement me")
}

func (M MHttpModule) ServeHTTP(writer http.ResponseWriter, request *http.Request) error {
	//TODO implement me
	panic("implement me")
}

func (M MHttpModule) Next() gateway.Module {
	//TODO implement me
	panic("implement me")
}
