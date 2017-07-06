package redis

//import (
//	"github.com/go-redis/redis"
//)
//
//type RedisDB struct {
//	DB  *redis.Client
//}
//
//func NewRedis(url, password string, db int) (redisDB *RedisDB, err error){
//	client := redis.NewClient(&redis.Options{
//		Addr:url,
//		Password:password,
//		DB:db,
//	})
//	redisDB = &RedisDB{DB:client}
//
//	return
//}
/* 创建reids客户端，具体方法由开发人员添加
func(p *RedisDB)Set(key, value string) bool{
	err := p.DB.Set(key, value, 0).Err()
	if err != nil{
		log.Fatal(err)
		return false
	}else {
		return true
	}
}
*/
