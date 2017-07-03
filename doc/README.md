#开发流程
1，go版本
1.8.3

2，IDE推荐
gogland

3，目录结构


3，日志库
日志库：github.com/cihub/seelog
介绍：godoc.org/github.com/cihub/seelog

//如果docker部署可以通过jenkins挂载配置文件进去
//否则使用viper或者confd
4，配置文件
使用:toml
介绍：godoc.org/github.com/BurntSushi/toml

5，依赖包管理-govendor
使用：go get github.com/kardianos/govendor
介绍：https://github.com/kardianos/govendor
     godoc.org/github.com/kardianos/govendor
     
6，单元测试工具-gotests
使用：go get github.com/cweill/gotests
介绍：https://github.com/cweill/gotests
     https://godoc.org/github.com/cweill/gotests
