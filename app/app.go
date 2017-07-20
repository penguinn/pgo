package app

import (
	"html/template"
	"github.com/penguinn/pgo/container"
	"github.com/penguinn/pgo/database/mysql"
	"gopkg.in/mgo.v2"
	"github.com/penguinn/pgo/database/mongo"
	"github.com/penguinn/pgo/database/redis"
)

func Register(name string, creator container.Creator) {
	container.DefaultContainer.Register(name, creator)
}

func ConfigureAll(cfg map[string]interface{}) error {
	return container.DefaultContainer.ConfigureAll(cfg)
}

func Get(names ...string) (interface{}, error) {
	return container.DefaultContainer.Get(names...)
}

func GetContainer(name string) (*container.Container, error) {
	return container.DefaultContainer.GetContainer(name)
}

//func GetLogger(name string) *log.Logger {
//	l, _ := defaultContainer.Get("logger", name)
//	ll, _ := l.(*log.Logger)
//	return ll
//}

//func GetRouter() *routing.Router {
//	r, _ := defaultContainer.Get("router")
//	rr, _ := r.(*routing.Router)
//	return rr
//}

func GetMongo(name string) (*mgo.Database, error){
	db, err := container.DefaultContainer.Get("mongo", name)
	if err == nil {
		if mDB, ok := db.(*mgo.Database); ok {
			return mDB, nil
		}
		return nil, err
	}
	return nil, err
}

//func GetRedis(name string) (*redis.Client, error) {
//	r, err := container.DefaultContainer.Get("redis", name)
//	if err == nil {
//		if rr, ok := r.(*redis.Client); ok {
//			return rr, nil
//		}
//		return nil, err
//	}
//	return nil, err
//}

func GetMySQL(name string) (*mysql.DB, error) {
	instance, err := container.DefaultContainer.Get("mysql", name)
	if err != nil {
		return nil, err
	}
	if db1, ok := instance.(*mysql.DB); ok {
		return db1, nil
	}
	return nil, err
}

func GetTemplate() *template.Template {
	t, _ := container.DefaultContainer.Get("template")
	tpl, _ := t.(*template.Template)
	return tpl
}

type Model interface {
	ConnName() string
}

func UseModel(name string, m Model, write bool) interface{} {
	d, err := Get(name, m.ConnName())
	if err == nil {
		switch name {
		case "mysql":
			return d.(*mysql.DB).Get(write)
		case "mongo":
			return d.(*mongo.MongoDB).Get(write)
		case "redis":
			return d.(*redis.RedisDB).Get(write)
		}
	}
	return nil
}