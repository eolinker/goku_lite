package redis

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

var (
	redisHost     string
	redisDB       int
	redisPort     string
	redisPassword string
	redisMode     string
	err           error
)

var option *redis.Options = &redis.Options{
	Addr:     "47.106.107.73:6379",
	Password: "",
	DB:       1,
	PoolSize: 2000,
}

var sentinelOption *redis.Options = &redis.Options{
	Addr:     "47.106.107.73:26379",
	Password: "",
	// DB:       1,
	PoolSize: 2000,
}

var failoverOptions *redis.FailoverOptions = &redis.FailoverOptions{
	MasterName: "mymaster3",
	// A seed list of host:port addresses of sentinel nodes.
	SentinelAddrs: []string{"47.106.107.73:26379", "47.106.107.73:26380", "47.106.107.73:26381"},
	// Following options are copied from Options struct.
	Password: "",
	DB:       0,
	PoolSize: 2000,
}

var clusterOptions *redis.ClusterOptions = &redis.ClusterOptions{
	ReadOnly:       true,
	RouteByLatency: true,
	Addrs:          []string{"112.74.38.5:7000", "112.74.38.5:7001", "112.74.38.5:7002", "119.23.210.62:7003", "119.23.210.62:7004", "119.23.210.62:7005"},
	Password:       "",
	// PoolSize:       2000,
	// DialTimeout:    10 * time.Second,
	// MinIdleConns:   10,
}

var ringOptions *redis.RingOptions = &redis.RingOptions{
	Addrs: map[string]string{
		"mymaster": "47.106.107.73:26379",
	},
	Password: "",
	DB:       0,
	PoolSize: 2000,
}

func sentinelMode() {
	t1 := time.Now()
	client := redis.NewFailoverClient(failoverOptions)
	fmt.Println(client.Info())
	fmt.Println("sentinel execute time:", time.Now().Sub(t1))
}

func sentinel() {
	t1 := time.Now()
	client := redis.NewSentinelClient(sentinelOption)
	fmt.Println(client.Sentinels("mymaster"))
	fmt.Println("sentinel execute time:", time.Now().Sub(t1))
}

func sentinelRingModeTest() {
	t1 := time.Now()
	sentinelAddrs := map[string][]string{
		"mymaster":  []string{"47.106.107.73:26379", "47.106.107.73:26380", "47.106.107.73:26381"},
		"mymaster2": []string{"47.106.107.73:26379", "47.106.107.73:26380", "47.106.107.73:26381"},
		"mymaster3": []string{"47.106.107.73:26379", "47.106.107.73:26380", "47.106.107.73:26381"},
		"mymaster4": []string{"47.106.107.73:26379", "47.106.107.73:26380", "47.106.107.73:26381"},
	}
	masterAddrMap := map[string]string{}
	for masterName, sentinelAddr := range sentinelAddrs {
		for _, addr := range sentinelAddr {
			sentinel := redis.NewSentinelClient(&redis.Options{
				Addr:     addr,
				PoolSize: 2000,
			})
			masterAddr, err := sentinel.GetMasterAddrByName(masterName).Result()
			if err != nil {
				_ = sentinel.Close()
				continue
			}
			masterAddrMap[masterName] = net.JoinHostPort(masterAddr[0], masterAddr[1])
			break
		}
	}
	sentinelRingOption := &redis.RingOptions{
		Addrs:    masterAddrMap,
		Password: "",
		DB:       0,
		PoolSize: 2000,
	}
	ringClient := redis.NewRing(sentinelRingOption)
	// ringClient.
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		err := ringClient.ForEachShard(func(c *redis.Client) error {
			// fmt.Println(c.Info())
			fmt.Println(c.String())
			return nil
		})
		if err != nil {
			panic(err)
		}
	}

	// fmt.Println(ringClient.Info())
	fmt.Println("sentinel execute time:", time.Now().Sub(t1))
}

func sentinelRingTest() {
	//sentinelAddrs := map[string][]string{
	//	"mymaster":  []string{"47.106.107.73:26379", "47.106.107.73:26380", "47.106.107.73:26381"},
	//	"mymaster2": []string{"47.106.107.73:26379", "47.106.107.73:26380", "47.106.107.73:26381"},
	//	"mymaster3": []string{"47.106.107.73:26379", "47.106.107.73:26380", "47.106.107.73:26381"},
	//	"mymaster4": ,
	//}
	addrs := []string{"47.106.107.73:26379", "47.106.107.73:26380", "47.106.107.73:26381"}
	masters := []string{"mymaster", "mymaster2", "mymaster3", "mymaster4"}
	sentinelRingOption := &SentinelRingOptions{
		Addrs:    addrs,
		Masters:  masters,
		Password: "",
		DB:       0,
		PoolSize: 2000,
	}
	sentinelRingClient := NewSentinelRing(sentinelRingOption)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		err := sentinelRingClient.ForEachAddr(func(addr string) error {
			fmt.Println(addr)
			// info, _ := c.Info("Server").Result()
			// fmt.Println(matchRedisInfo(info, "tcp_port"))
			return nil
		})
		fmt.Println("--------------------------------------------")
		if err != nil {
			panic(err)
		}
	}
}

func matchRedisInfo(ret string, dest string) string {
	retArr := strings.Split(strings.TrimRight(ret, "\r\n"), "\r\n")
	retMap := make(map[string]string)

	dataBases := make([]map[string]string, 0)

	for _, v := range retArr {
		if index := strings.Index(v, "#"); index == -1 && v != "" {

			kvArr := strings.Split(v, ":")
			k, v := kvArr[0], kvArr[1]
			retMap[k] = v
			if strings.HasPrefix(k, "db") == true {
				vArr := strings.Split(v, ",")
				keyArr := strings.Split(vArr[0], "=")
				keys := keyArr[1]
				expiresArr := strings.Split(vArr[1], "=")
				expires := expiresArr[1]
				database := map[string]string{
					"name":    k,
					"keys":    keys,
					"expires": expires,
				}
				dataBases = append(dataBases, database)
			}

		}
	}
	if v, ok := retMap[dest]; ok {
		return v
	}
	return ""
}

func sentinelRingNode() {
	t1 := time.Now()
	client := redis.NewSentinelClient(sentinelOption)
	// ipList := make([]string, 0)
	// ipList, _ = client.GetMasterAddrByName("mymaster").Result()
	fmt.Println(client.GetMasterAddrByName("mymaster").Result())
	fmt.Println("sentinel execute time:", time.Now().Sub(t1))
}

func ringMode() {
	t1 := time.Now()
	ringClient := redis.NewRing(ringOptions)
	ringClient.ForEachShard(func(c *redis.Client) error {
		fmt.Println(c.Info())
		return nil
	})

	fmt.Println("ring execute time:", time.Now().Sub(t1))
}

func main() {
	// sentinelMode()
	sentinelRingTest()
	// ringMode()
	// //建立连接
	// sf := &redis.FailoverOptions{
	// 	// The master name.
	// 	MasterName: "mymaster",
	// 	// A seed list of host:port addresses of sentinel nodes.
	// 	SentinelAddrs: []string{"47.106.107.73:26379", "47.106.107.73:26380", "47.106.107.73:26381"},

	// 	// Following options are copied from Options struct.
	// 	Password: "",
	// 	DB:       0,
	// }
	// client := redis.NewFailoverClient(sf)

	// // fmt.Println("client:", clusterClient)
	// t1 := time.Now()
	// t2 := time.Now()
	// clusterClient := redis.NewClusterClient(clusterOptions)
	// for i := 0; i < 10; i++ {
	// 	// res, err1 := clusterClient.Ping().Result()
	// 	// res, err1 := clusterClient.Do("ping").Result()
	// 	fmt.Println("Hits1:", clusterClient.PoolStats().Hits)
	// 	fmt.Println("IdleConns1:", clusterClient.PoolStats().IdleConns)
	// 	fmt.Println("Misses1:", clusterClient.PoolStats().Misses)
	// 	fmt.Println("StaleConns1:", clusterClient.PoolStats().StaleConns)
	// 	fmt.Println("Timeouts1:", clusterClient.PoolStats().Timeouts)
	// 	fmt.Println("TotalConns1:", clusterClient.PoolStats().TotalConns)
	// 	res, err1 := clusterClient.Keys("goku*").Result()
	// 	// cmd := clusterClient.Keys("goku*")
	// 	// fmt.Println("cmd:", cmd)
	// 	// fmt.Println("Ping:", res, err1)
	// 	fmt.Println("Len1:", len(res), err1)
	// 	keys := make([]string, 0)
	// 	clusterClient.ForEachNode(func(client *redis.Client) error {
	// 		keyList, _ := client.Keys("goku*").Result()
	// 		keys = append(keys, keyList...)
	// 		return nil
	// 	})

	// 	// cmd := clusterClient.Keys("goku*")
	// 	// fmt.Println("cmd:", cmd)
	// 	// fmt.Println("Ping:", res, err1)
	// 	fmt.Println("Len2:", len(keys), err1)
	// 	// fmt.Println("status:", clusterClient.PoolStats())
	// 	fmt.Println("Hits:", clusterClient.PoolStats().Hits)
	// 	fmt.Println("IdleConns:", clusterClient.PoolStats().IdleConns)
	// 	fmt.Println("Misses:", clusterClient.PoolStats().Misses)
	// 	fmt.Println("StaleConns:", clusterClient.PoolStats().StaleConns)
	// 	fmt.Println("Timeouts:", clusterClient.PoolStats().Timeouts)
	// 	fmt.Println("TotalConns:", clusterClient.PoolStats().TotalConns)
	// 	if i == 0 {
	// 		t2 = time.Now()
	// 	}
	// 	break
	// 	// clusterClient.Close()
	// }
	// // client := redis.NewClient(option)
	// // for i := 0; i < 10; i++ {
	// // 	// res, err1 := client.Ping().Result()
	// // 	// res, err1 := client.Do("ping").Result()
	// // 	res, err1 := client.Keys("goku*").Result()
	// // 	// fmt.Println("Ping:", res, err1)
	// // 	fmt.Println("Keys:", res, err1)
	// // 	// fmt.Println("status:", clusterClient.PoolStats())
	// // 	fmt.Println("Hits:", client.PoolStats().Hits)
	// // 	fmt.Println("IdleConns:", client.PoolStats().IdleConns)
	// // 	fmt.Println("Misses:", client.PoolStats().Misses)
	// // 	fmt.Println("StaleConns:", client.PoolStats().StaleConns)
	// // 	fmt.Println("Timeouts:", client.PoolStats().Timeouts)
	// // 	fmt.Println("TotalConns:", client.PoolStats().TotalConns)
	// // 	// clusterClient.Close()
	// // 	if i == 0 {
	// // 		t2 = time.Now()
	// // 	}
	// // }

	// // client := redis.NewFailoverClient(failoverOptions)
	// // for i := 0; i < 10; i++ {
	// // 	// res, err1 := client.Ping().Result()
	// // 	// res, err1 := client.Do("ping").Result()
	// // 	res, err1 := client.Keys("goku*").Result()
	// // 	// fmt.Println("Ping:", res, err1)
	// // 	fmt.Println("Keys:", res, err1)
	// // 	// fmt.Println("status:", clusterClient.PoolStats())
	// // 	fmt.Println("Hits:", client.PoolStats().Hits)
	// // 	fmt.Println("IdleConns:", client.PoolStats().IdleConns)
	// // 	fmt.Println("Misses:", client.PoolStats().Misses)
	// // 	fmt.Println("StaleConns:", client.PoolStats().StaleConns)
	// // 	fmt.Println("Timeouts:", client.PoolStats().Timeouts)
	// // 	fmt.Println("TotalConns:", client.PoolStats().TotalConns)
	// // 	// clusterClient.Close()
	// // }
	// fmt.Println("DialTimeout:", clusterOptions.DialTimeout)
	// fmt.Println("waste time:", time.Now().Sub(t1))
	// fmt.Println("waste time t2:", time.Now().Sub(t2))

}
