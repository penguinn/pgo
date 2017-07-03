package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

//创建的数据库打开sql日志，并会自动创建连接池，连接池最多空闲10个连接，最多有100个连接

type MySQLDB struct {
	DB   	*gorm.DB
}

func NewMySQLDB(host, password, user, database string) (mysqlDB *MySQLDB, err error){
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, database)
	db, err := gorm.Open("mysql", args)
	if err != nil{
		return nil, err
	}
	//db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	mysqlDB = &MySQLDB{DB:db}
	return
}

/*
此文件只负责创建数据库，表结构和orm函数写在其他文件，example:
type UserDevice struct {
	ID 				uint 	`gorm:"column:id;primary_key;not null;unsigned data type;AUTO_INCREMENT"`
	UserID			uint	`gorm:"column:userId;not null;unsigned data type"`
	DeviceNum		string	`gorm:"column:deviceNum;type:varchar(64);not null;DEFAULT:''"`
	InsertTime		uint64	`gorm:"column:insertTime;not null;unsigned data type"`
	ModifyTime		uint64	`gorm:"column:modifyTime;not null;unsigned data type"`
	Platform		string	`gorm:"column:platform;type:char(1);not null;DEFAULT:''"`
	Source 			uint8  	`gorm:"column:source;type:tinyint(3);not null;unsigned data type;DEFAULT:1"`
	IsOff			bool	`gorm:"column:isOff;type:tinyint(1);not null;DEFAULT:0"`
	ExtId 			string	`gorm:"column:extId;type:varchar(128);not null;DEFAULT:''"`
	PushPlatform	string	`gorm:"column:pushPlatform;type varchar(20);not null"`
	Version        int64 	`gorm:"column:version;not null"`
}

func(p *MySQLDB) InsertOneFromUserDevice(userID int, deviceNum string, devicePlatform string, source int, extId string,
		pushPlatform string, version int64) error {
	query := UserDevice{UserID:uint(userID), DeviceNum:deviceNum, InsertTime:uint64(time.Now().Unix()),
		ModifyTime:uint64(time.Now().Unix()), Platform:devicePlatform, Source:uint8(source), ExtId:extId, PushPlatform:pushPlatform, Version:version}
	//这个地方会插入 然后select字段为空的数据保存到struct中
	p.db.Table("userDevice").Create(&query)
	//fmt.Println(query.ID)
	return p.db.Error
}

 */