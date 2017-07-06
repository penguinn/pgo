package mongo

//import (
//	"gopkg.in/mgo.v2"
//)
//
//type MongoDB struct {
//	db 			*mgo.Database
//}
//
//func NewMongoDB(url, username, password string) (db *MongoDB, err error){
//	session, err := mgo.Dial(url)
//	if err != nil{
//		return nil, err
//	}
//	session.Login(&mgo.Credential{Username:username, Password:password})
//	session.SetMode(mgo.Monotonic, true)
//	mdb = session.DB("test")
//  db := &MongoDB{db:mdb}
//	return
//}
/*
创建的类以数据库为单位，example:
type Hehe struct {
	Id 		float32
	Value 	float32
}

func (p *MongoDB) GetValueFromHehebyId(collection string, id float32) (value float32, err error){
	result := Hehe{}
	err = p.db.C(collection).Find(bson.M{"id":id}).One(result)
	if err != nil{
		return
	}else{
		value = result.Value
		return
	}
}
*/
