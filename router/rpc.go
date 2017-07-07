package router

import (
	"github.com/penguinn/pgo/jsonrpc2"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type RouterRpcCfg struct {
	Path 		string
}

func InitRpc(controllers interface{}) {
	for _, controller := range controllers.([]jsonrpc2.Controller){
		jsonrpc2.RegisterMethod(controller.Name, controller.F, controller.Params, controller.Result)
	}
}

func NewRpcRouter(cfg *RouterRpcCfg) {
	http.HandleFunc(cfg.Path, jsonrpc2.Handler)
}

func CreatorRpc(cfg interface{}) (interface{}, error) {
	cfgMap := cfg.(map[string]interface {})["default"]
	for _, cfgTmp := range cfgMap.([]interface{}){
		var routerCfg RouterRpcCfg
		mapstructure.WeakDecode(cfgTmp, &routerCfg)
		NewRpcRouter(&routerCfg)
	}
	return nil, nil
}
