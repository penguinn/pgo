package pgo

import (
	"github.com/spf13/viper"
	"log"
	"github.com/penguinn/pgo/app"
	"github.com/penguinn/pgo/database/mysql"
	"github.com/penguinn/pgo/database/redis"
	"github.com/penguinn/pgo/router"
	"errors"
	"net/http"
	"github.com/penguinn/pgo/template"
	"github.com/penguinn/pgo/log"
	"github.com/penguinn/pgo/thrift"
	"github.com/penguinn/pgo/database/mongo"
	"github.com/penguinn/pgo/prpcx"
)

func Init(configFile string, args ...interface{}) error {
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil{
		errStr := "Config parse error," + err.Error()
		log.Fatal(errStr)
		return err
	}else {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}

	//初始化日志
	pLog.Init(viper.GetString("server.log"))

	//初始化服务类型
	runMode := viper.GetString("server.type")
	if len(runMode) != 0{
		switch runMode {
		case "http":
			if len(args) == 1{
				router.InitHttp(args[0])
			}else{
				return errors.New("args error")
			}
		case "thrift":
			if len(args) == 1 {
				thrift.Init(args[0])
			} else{
				return errors.New("args error")
			}
		case "jsonrpc2":
			if len(args) == 1 {
				router.InitRpc(args[0])
			} else{
				return errors.New("args error")
			}
		case "rpcx":
			if len(args) == 2{
				prpcx.InitRpcx(args[0].(string), args[1])
			} else {
				return errors.New("args error")
			}
		default:
			return errors.New("run error with type："+runMode)
		}
	}

	//初始化组件
	if len(viper.GetStringMap("components.mysql")) != 0 {
		app.Register("mysql", mysql.Creator)
	}
	if len(viper.GetStringMap("components.mongo")) != 0{
		app.Register("mongo", mongo.Creator)
	}
	if len(viper.GetStringMap("components.redis")) != 0 {
		app.Register("redis", redis.Creator)
	}
	if len(viper.GetStringMap("components.router")) != 0 && viper.GetString("components.router.type") == "web"{
		app.Register("http", router.CreatorHttp)
	}
	if len(viper.GetStringMap("components.router")) != 0 && viper.GetString("components.router.type") == "jsonrpc2"{
		app.Register("rpc", router.CreatorRpc)
	}
	if len(viper.GetStringMap("components.template")) != 0 {
		app.Register("template", template.Creator)
	}
	//fmt.Println(viper.GetStringMap("components.mysql"))
	return app.ConfigureAll(viper.GetStringMap("components"))
}

func Run() error {
	runMode := viper.GetString("server.type")
	if len(runMode) != 0{
		switch runMode {
		case "http":
			err := http.ListenAndServe(viper.GetString("server.addr"), nil)
			if err != nil{
				log.Fatal("ListenAndServer: ", err)
			}
		case "thrift":
			err := thrift.Run()
			if err != nil{
				log.Fatal("RunThriftServer: ", err)
			}
		case "jsonrpc2":
			err := http.ListenAndServe(viper.GetString("server.addr"), nil)
			if err != nil{
				log.Fatal("ListenAndServer: ", err)
			}
		case "rpcx":
			err := prpcx.Run(viper.GetString("server.addr"))
			if err != nil{
				log.Fatal("RunRpcxServer: ", err)
			}
		default:
			return errors.New("run error with type："+runMode)
		}
	}
	return errors.New("run error with type："+runMode)
}