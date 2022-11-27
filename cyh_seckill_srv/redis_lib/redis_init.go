package redis_lib

import (
	"github.com/gomodule/redigo/redis"
)

var Conn redis.Conn
var Err error

func InitRedis() {
	Conn, Err = redis.Dial("tcp", "127.0.0.1:6379")
	if Err != nil {
		panic(Err)
	}
}

func CloseRedis() {
	Err = Conn.Close()
	if Err != nil {
		panic(Err)
	}
}
