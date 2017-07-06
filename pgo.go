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
)

func Init(configFile string, controller interface{}) error {
	if controller != nil{
		router.Init(controller)
	}
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

	if len(viper.GetStringMap("components.mysql")) != 0 {
		app.Register("mysql", mysql.Creator)
	}
	if len(viper.GetStringMap("components.redis")) != 0 {
		app.Register("redis", redis.Creator)
	}
	if len(viper.GetStringMap("components.router")) != 0 {
		app.Register("router", router.Creator)
	}
	//fmt.Println(viper.GetStringMap("components.mysql"))
	return app.ConfigureAll(viper.GetStringMap("components"))
}

func Run() error {
	runMode := viper.GetString("server.type")
	if len(runMode) != 0{
		switch runMode {
		case "web":
			err := http.ListenAndServe(viper.GetString("server.addr"), nil)
			if err != nil{
				log.Fatal("ListenAndServer: ", err)
			}
		case "thrift":
			//todo
		case "jsonrpc2":
			//todo
		default:
			return errors.New("run error with type："+runMode)
		}
	}
	return errors.New("run error with type："+runMode)
}