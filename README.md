# pgo #
* go的基础开发框架，在这里集成了日志，配置文件，数据库等东西，目前仅支持web，thrift，jsonrpc2的开发 

## 开发流程 ##
1. go版本
* 1.8.3

2. IDE推荐
* gogland

3. 目录结构<br />
![Alt text](./doc/construction.png "结构图片")
* 第一个红框表示GOPATH的路径
* 第二个红框是go get下来生成的二进制文件保存的地方，最好把这个路径添加到PATH
* 第三个红框是go get下载的github项目保存的目录
* 第四个红框是我们自己的gitlab的地址，我们以后的项目都在这个目录下开发（此图中MedPush就是一个项目）
* 第五个红框是govendor生成的包管理目录
* 第六个是没个项目的Dockerfile，用于docker build

4. 日志库<br />
* 日志库：github.com/cihub/seelog
* 介绍：https://godoc.org/github.com/cihub/seelog

<!-- 如果docker部署可以通过jenkins挂载配置文件进去
否则使用viper或者confd -->
5. 配置文件
* 使用：toml
* 介绍：https://godoc.org/github.com/BurntSushi/toml
* 说明：如果docker部署可以通过jenkins挂载配置文件进去，否则使用viper或者confd

6. 依赖包管理-govendor
* 使用：go get github.com/kardianos/govendor
* 介绍：https://github.com/kardianos/govendor  
　　　https://godoc.org/github.com/kardianos/govendor
     
7. 单元测试工具-gotests
* 使用：go get github.com/cweill/gotests
* 介绍：https://github.com/cweill/gotests  
　　　https://godoc.org/github.com/cweill/gotests

8. 数据库介绍  
* orm:  https://godoc.org/github.com/jinzhu/gorm
* mongo: https://godoc.org/gopkg.in/mgo.v2
* redis: https://godoc.org/github.com/go-redis/redis

## 开发实例 ##
1. web开发  
* 位置：https://github.com/penguinn/pgo-test/tree/master/web
* 注意：初始化的时候需要传入handler实例的地址

2. thrift开发 
* 位置：https://github.com/penguinn/pgo-test/tree/master/RPCThrift  
* 注意：初始化的时候需要传入通过handler生成的processor

3. 基于http的jsonrpc2开发
* 位置：https://github.com/penguinn/pgo-test/tree/master/JSONRPC2
* 注意：初始化的时候需要传入需要注册的类
