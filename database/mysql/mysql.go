package mysql

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"

	"github.com/mitchellh/mapstructure"
	"fmt"
)

type DB struct {
	Config *DBConfig
	Write  *gorm.DB
	Reads  []*gorm.DB
}

type DBConfig struct {
	Driver string
	Dsn    string
	Reads  []string
}

func NewDB(cfg *DBConfig) (*DB, error) {
	var err error
	db := &DB{
		Config: cfg,
	}
	db.Write, err = gorm.Open(cfg.Driver, cfg.Dsn)
	if err != nil {
		return nil, err
	}
	for dsn := range cfg.Reads {
		rdb, err := gorm.Open(cfg.Driver, dsn)
		if err != nil {
			return nil, err
		}
		db.Reads = append(db.Reads, rdb)
	}
	return db, nil
}

func (this *DB) Get(write bool) *gorm.DB {
	if write {
		return this.Write
	}
	l := len(this.Reads)
	if l == 0 {
		return this.Write
	}
	rand.Seed(time.Now().UnixNano())
	return this.Reads[rand.Intn(l)]

}

func Creator(cfg interface{}) (interface{}, error) {
	fmt.Println(1)
	var dbConfig DBConfig
	err := mapstructure.WeakDecode(cfg, &dbConfig)
	if err != nil {
		return nil, err
	}
	return NewDB(&dbConfig)
}
