package goku_plugin

var (
	_redisManager RedisManager
	_logger Logger
)
func InitLog(logger Logger){
	_logger = logger
}
func SetRedisManager(manager RedisManager) {
	if _redisManager != nil {
		panic("repeat set RedisManager")
	}
	_redisManager = manager
}


func GetRedis() Redis {
	if _redisManager == nil {
		return nil
	}
	return _redisManager.Default()
}
func GetRedisByName(name string) (Redis, bool) {
	return _redisManager.Get(name)
}

const version  = "20190808"
func Version()  string {
	return version
}