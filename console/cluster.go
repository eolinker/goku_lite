package console

import (
	"io/ioutil"
	"strconv"

	"github.com/eolinker/goku/common/conf"
	"github.com/eolinker/goku/common/redis-manager"
	cluster2 "github.com/eolinker/goku/server/cluster"
	console_mysql "github.com/eolinker/goku/server/dao/console-mysql"
	"github.com/eolinker/goku/server/entity"
	"gopkg.in/yaml.v2"
)

func loadCluster(file string) ([]*entity.ClusterInfo, error) {

	body, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	tv := struct {
		Cluster []*entity.ClusterInfo `yaml:cluster`
	}{}

	err = yaml.Unmarshal(body, &tv)

	if err != nil {
		return nil, err
	}
	cs := make([]*entity.ClusterInfo, 0, len(tv.Cluster)+1)
	for _, v := range tv.Cluster {
		cs = append(cs, v)
	}

	return cs, nil
}
func getDefaultDatabase() (*entity.ClusterDB, error) {
	dbPort, err := strconv.Atoi(conf.MastValue("db_port", "3306"))
	if err != nil {
		return nil, err
	}
	return &entity.ClusterDB{
		Driver:   conf.MastValue("db_type", "mysql"),
		Host:     conf.Value("db_host"),
		Port:     dbPort,
		UserName: conf.Value("db_user"),
		Password: conf.Value("db_password"),
		Database: conf.Value("db_name"),
	}, nil
}
func InitClusters() {

	infos, err := loadCluster(conf.MastValue("cluster_config", "config/cluster.yaml"))
	if err != nil {
		panic(err)
	}
	all := infos

	err = console_mysql.InsertClusters(all)
	if err != nil {
		panic(err)
	}
	cs, err := console_mysql.LoadClusters()

	if err != nil {
		panic(err)
	}

	currentClusters := make([]*entity.ClusterInfo, 0, len(infos))
	redisOfCluster := make(map[string]redis_manager.RedisConfig)

	for _, info := range infos {
		c := cs[info.Name]
		currentClusters = append(currentClusters, c)
		redisOfCluster[c.Name] = c.Redis
	}

	redis_manager.InitRedisOfCluster(redisOfCluster)

	cluster2.Init(currentClusters)

}
