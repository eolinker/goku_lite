package monitor_read

import (
	"fmt"
	redis_manager "github.com/eolinker/goku/common/redis-manager"
	"github.com/eolinker/goku/server/dao/console-mysql/dao-monitor"
	"github.com/eolinker/goku/server/entity"
	"github.com/eolinker/goku/server/monitor/monitor-key"
	"strconv"
	"time"
)
var (
 	period = 30*time.Second
)

func SetPeriod(sec int)  {
	period = time.Duration(sec)* time.Second
}
func InitMonitorRead(clusters []*entity.Cluster) error {
	for _,c:=range clusters{
		_,has:= redis_manager.Get(c.Name)
		if !has{
			return fmt.Errorf("no redis for cluster:%s",c.Name)
		}
	}
	for _,c:=range clusters{
		go doLoopForCluster(c.Name,c.Id)
	}
	return nil
}

func doLoopForCluster(clusterName string,clusterId int)  {


	t:=time.NewTimer(period)
	
	for{
		select {
		case <-t.C:{
			read(clusterName,clusterId,time.Now())
		}
		}
		t.Reset(period)
	}

}
func read(clusterName string,clusterId int,t time.Time)  {
	hour:= t.Format("2006010215")
	now:= t.Format("2006-01-02 15:04:05")


	hourValue ,_:=strconv.Atoi(hour)

	// 包含 strate == ""
	strategyIds,err:=readStrategyId(hour,clusterName)
	if err!= nil{
		return
	}
	for _,strategyId:=range strategyIds{

		apiIds,err:=readApiId(hour,clusterName,strategyId)
		if err!= nil{
			continue
		}

		for _,apiId:=range apiIds{

			valus,err:=readValue(hour,clusterName,strategyId,apiId)
			if err!=nil{
				continue
			}
			apiID ,_:= strconv.Atoi(apiId)

			dao_monitor.Save(strategyId,apiID,clusterId,hourValue,now,valus)
		}
	}

}
func readStrategyId(now,cluster string)([]string ,error) {
	key:= monitor_key.StrategyMapKey(cluster,now)
	conn,_:= redis_manager.Get(cluster)
	return conn.HKeys(key).Result()

}
func readApiId(now ,cluster,strategyId string)([]string ,error){
	key:= monitor_key.APiMapKey(cluster,strategyId,now)
	conn,_:= redis_manager.Get(cluster)
	return conn.HKeys(key).Result()
}
func readValue(now,cluster,strategyId string,apiId string) (map[string]string,error) {
	conn,_:= redis_manager.Get(cluster)
	key := monitor_key.ApiValueKey(cluster,strategyId,apiId,now)
	return conn.HGetAll(key).Result()

}