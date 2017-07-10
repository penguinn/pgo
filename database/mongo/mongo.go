package mongo

import (
	"gopkg.in/mgo.v2"
	"github.com/mitchellh/mapstructure"
	"log"
	"time"
)

type MongoDB struct {
	Config 		*MongoConfig
	Mgo 		*mgo.Database
}

type MongoConfig struct {
	Addresses 		[]string
	UserName 		string
	Password 		string
	Database		string
}

func NewMongoDB(cfg *MongoConfig) (db *MongoDB, err error){
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:cfg.Addresses,
		Username:cfg.UserName,
		Password:cfg.Password,
		Timeout:2*time.Second,
	})
	db = &MongoDB{
		Config:cfg,
	}
	if err != nil{
		log.Fatal("New mongo error:", err)
		return
	}
	session.SetMode(mgo.Monotonic, true)
	db.Mgo = session.DB(cfg.Database)
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

func (this *MongoDB) Get(write bool) *mgo.Database {

	return this.Mgo

}
