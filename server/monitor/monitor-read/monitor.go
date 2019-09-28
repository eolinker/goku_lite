package monitorread

import (
	"fmt"
	redis_manager "github.com/eolinker/goku-api-gateway/common/redis-manager"
	"github.com/eolinker/goku-api-gateway/server/dao/console-mysql/dao-monitor"
	"github.com/eolinker/goku-api-gateway/server/entity"
	"github.com/eolinker/goku-api-gateway/server/monitor/monitor-key"
	"strconv"
	"time"
)

var (
	period = 30 * time.Second
)

//SetPeriod 设置更新周期
func SetPeriod(sec int) {
	period = time.Duration(sec) * time.Second
}

//InitMonitorRead init monitor read
func InitMonitorRead(clusters []*entity.Cluster) error {
	for _, c := range clusters {
		_, has := redis_manager.Get(c.Name)
		if !has {
			return fmt.Errorf("no redis for cluster:%s", c.Name)
		}
	}
	for _, c := range clusters {
		go doLoopForCluster(c.Name, c.ID)
	}
	return nil
}

func doLoopForCluster(clusterName string, clusterID int) {

	t := time.NewTimer(period)

	for {
		select {
		case <-t.C:
			{
				read(clusterName, clusterID, time.Now())
			}
		}
		t.Reset(period)
	}

}
func read(clusterName string, clusterID int, t time.Time) {
	hour := t.Format("2006010215")
	now := t.Format("2006-01-02 15:04:05")

	hourValue, _ := strconv.Atoi(hour)

	// 包含 strate == ""
	strategyIds, err := readStrategyID(hour, clusterName)
	if err != nil {
		return
	}
	for _, strategyID := range strategyIds {

		apiIds, err := readAPIId(hour, clusterName, strategyID)
		if err != nil {
			continue
		}

		for _, apiID := range apiIds {

			valus, err := readValue(hour, clusterName, strategyID, apiID)
			if err != nil {
				continue
			}
			apiID, _ := strconv.Atoi(apiID)

			dao_monitor.Save(strategyID, apiID, clusterID, hourValue, now, valus)
		}
	}

}
func readStrategyID(now, cluster string) ([]string, error) {
	key := monitorkey.StrategyMapKey(cluster, now)
	conn, _ := redis_manager.Get(cluster)
	return conn.HKeys(key).Result()

}
func readAPIId(now, cluster, strategyID string) ([]string, error) {
	key := monitorkey.APIMapKey(cluster, strategyID, now)
	conn, _ := redis_manager.Get(cluster)
	return conn.HKeys(key).Result()
}
func readValue(now, cluster, strategyID string, apiID string) (map[string]string, error) {
	conn, _ := redis_manager.Get(cluster)
	key := monitorkey.APIValueKey(cluster, strategyID, apiID, now)
	return conn.HGetAll(key).Result()
}
