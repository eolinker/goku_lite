package redis_manager

//
//func GetKeys(r Redis, pattern string) ([]string, error) {
//	keys := []string{}
//	if redisMode == "stand-alone" {
//		r.(*localRedis.Ring).ForEachShard(func(c *localRedis.Client) error {
//			keySlice, _ := c.Keys(pattern).Result()
//			keys = append(keys, keySlice...)
//			return nil
//		})
//	} else if redisMode == "sentinel" {
//		r.(*localRedis.SentinelRing).ForEachShard(func(c *localRedis.Client) error {
//			keySlice, _ := c.Keys(pattern).Result()
//			keys = append(keys, keySlice...)
//			return nil
//		})
//	} else if redisMode == "cluster" {
//		r.(*localRedis.ClusterClient).ForEachNode(func(c *localRedis.Client) error {
//			keySlice, _ := c.Keys(pattern).Result()
//			keys = append(keys, keySlice...)
//			return nil
//		})
//	}
//	return keys, nil
//}
//
//func GetRedisMode() string {
//	return redisMode
//}
//
//func GetRingOption() *localRedis.RingOptions {
//	return ringClient.Options()
//}
//
//func GetSentinelRingOption() *localRedis.SentinelRingOptions {
//	return sentinelClient.Options()
//}
//
//func GetClusterOption() *localRedis.ClusterOptions {
//	return clusterClient.Options()
//}
