package router

import (
	"github.com/mitchellh/mapstructure"
	//"github.com/penguinn/pgo-test/web/handler"
	"reflect"
	"fmt"
	"net/http"
	"errors"
)

type ControllerMapsType map[string]reflect.Value

var controllerMaps = make(ControllerMapsType)

type RouterCfg struct {
	Path 		string
	Handler 	string
}

func Init(controller interface{}) {
	//var controller handler.Controller
	//controllerMap := make(ControllerMapsType, 0)
	//crMap := make(ControllerMapsType, 0)
	//创建反射变量，注意这里需要传入ruTest变量的地址；
	//不传入地址就只能反射Routers静态定义的方法
	vf := reflect.ValueOf(controller)
	fmt.Println(vf.Kind())
	vft := vf.Type()
	//读取方法数量
	mNum := vf.NumMethod()
	//遍历路由器的方法，并将其存入控制器映射变量中
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		fmt.Println("index:", i, " MethodName:", mName)
		controllerMaps[mName] = vf.Method(i) //<<<
	}
}

func NewRouter(cfg *RouterCfg) {
	container, ok := controllerMaps[cfg.Handler]
	if ok{
		handlerFunc := container.Interface().(func(http.ResponseWriter, *http.Request))
		http.HandleFunc(cfg.Path, handlerFunc)
	}else {
		panic(errors.New("have not this controller"))
	}

}

func Creator(cfg interface{}) (interface{}, error) {
	cfgMap := cfg.(map[string]interface {})["default"]
	for _, cfgTmp := range cfgMap.([]interface{}){
		var routerCfg RouterCfg
		mapstructure.WeakDecode(cfgTmp, &routerCfg)
		NewRouter(&routerCfg)
	}
	return nil, nil
}
