package cache

import (
	"goku-ce-1.0/conf"
	_ "goku-ce-1.0/utils"
	"strconv"

	"github.com/codegangsta/inject"
	"github.com/farseer810/yawf"
	redis "github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool
)

func init() {
	pool = redis.NewPool(dial(), 2000)
}

func dial() func() (redis.Conn, error) {
	db, err := strconv.Atoi(conf.Configure["redis_db"])
	if err != nil {
		db = 0
	}
	if _, hasPassword := conf.Configure["redis_password"]; hasPassword {
		return func() (redis.Conn, error) {
			return redis.Dial("tcp",
				conf.Configure["redis_host"]+":"+conf.Configure["redis_port"],
				redis.DialPassword(conf.Configure["redis_password"]),
				redis.DialDatabase(db))
		}
	} else {
		return func() (redis.Conn, error) {
			return redis.Dial("tcp",
				conf.Configure["redis_host"]+":"+conf.Configure["redis_port"],
				redis.DialDatabase(db))
		}
	}
}

func getConnection() redis.Conn {
	return pool.Get()
}

func GetConnectionFromContext(context yawf.Context) redis.Conn {
	var conn redis.Conn = nil
	value := context.Get(inject.InterfaceOf((*redis.Conn)(nil)))
	if value.IsValid() {
		conn = value.Interface().(redis.Conn)
	}
	return conn
}

func GetConnection(context yawf.Context) redis.Conn {
	conn := GetConnectionFromContext(context)
	if conn != nil {
		return conn
	}

	conn = getConnection()
	context.MapTo(conn, (*redis.Conn)(nil))
	return conn
}
