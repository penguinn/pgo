package mongo

import (
	"gopkg.in/mgo.v2"
	"github.com/mitchellh/mapstructure"
)

type MongoConfig struct {
	Addresses 		[]string
	UserName 		string
	Password 		string
}

func NewMongoDB(cfg *MongoConfig) (db *mgo.Database, err error){
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:cfg.Addresses,
		Username:cfg.UserName,
		Password:cfg.Password,
	})
	if err != nil{
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	db = session.DB("test")
	return
}

func Creator(cfg interface{}) (interface{}, error) {
	var mongoConfig MongoConfig
	err := mapstructure.WeakDecode(cfg, &mongoConfig)
	if err != nil{
		return nil, err
	}
	return NewMongoDB(&mongoConfig)
}
