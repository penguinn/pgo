package main

import(
	"github.com/spf13/viper"
)

func main() {
	var runtime_viper = viper.New()
	runtime_viper.AddRemoteProvider("etcd", "http:127.0.0.1:2379", "/config/example")
	runtime_viper.SetConfigType("toml")
}
