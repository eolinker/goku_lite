package static

import (
	"errors"
	"fmt"
	"time"

	"github.com/eolinker/goku-api-gateway/config"
	"github.com/eolinker/goku-api-gateway/goku-service/health"

	"strconv"
	"strings"
	"unicode"

	"github.com/eolinker/goku-api-gateway/goku-service/common"
)

//ErrorNoInstance errorNoInstance
var ErrorNoInstance = errors.New("no instance")

//Sources source
type Sources struct {
	name string

	discovery          *Discovery
	healthCheckHandler health.CheckHandler
	instanceFactory    *common.InstanceFactory
}

//SetHealthConfig setHealthConfig
func (s *Sources) SetHealthConfig(conf *config.HealthCheckConfig) {
	if conf == nil || !conf.IsHealthCheck {
		s.Close()
		return
	}

	s.healthCheckHandler.Open(
		conf.URL,
		conf.StatusCode,
		conf.Second,
		time.Duration(conf.TimeOutMill)*time.Millisecond)
}

//Close close
func (s *Sources) Close() {
	instances := s.healthCheckHandler.Close()
	for _, instance := range instances {
		instance.ChangeStatus(common.InstanceChecking, common.InstanceRun)
	}
}

//CheckDriver checkDriver
func (s *Sources) CheckDriver(driverName string) bool {

	return driverName == DriverName
}

//SetDriverConfig setDriverConfig
func (s *Sources) SetDriverConfig(config string) error {
	return nil
}

//GetApp getApp
func (s *Sources) GetApp(app string) (*common.Service, health.CheckHandler, bool) {
	service, e := s.decode(app)
	if e != nil {
		return nil, nil, false
	}
	return service, s.healthCheckHandler, true
}

//NewStaticSources 创建Sources
func NewStaticSources(name string) *Sources {
	return &Sources{

		name:               name,
		discovery:          new(Discovery),
		healthCheckHandler: &health.CheckBox{},
		instanceFactory:    common.NewInstanceFactory(),
	}
}

func fields(str string) []string {

	words := strings.FieldsFunc(strings.Join(strings.Split(str, ";"), " ; "), func(r rune) bool {
		if unicode.IsSpace(r) {
			return true
		}

		return false
	})
	return words
}

func (s *Sources) decode(config string) (*common.Service, error) {

	words := fields(config)

	instances := make([]*common.Instance, 0, 5)
	nodes := make([]*Node, 0, 5)
	var node *Node
	index := 0
	for _, word := range words {
		if word == ";" {
			index = 0
			node = nil
			continue
		}
		l := len(word)
		value := word
		if word[l-1] == ';' {
			value = word[:l-1]
		}
		switch index {
		case 0:
			{
				node = new(Node)
				vs := strings.Split(value, ":")
				if len(vs) > 2 {
					return nil, fmt.Errorf("decode ip:port failt for[%s]", value)
				}
				node.IP = vs[0]
				if len(vs) == 2 {
					node.Port, _ = strconv.Atoi(vs[1])
				}
				nodes = append(nodes, node)

			}
		case 1:
			{
				weight, err := strconv.Atoi(value)
				if err != nil {
					return nil, err
				}
				node.Weight = weight
			}
		case 2:
			{
				//node.Status = common.ParseStatus(value)
			}
		}
		if node.Weight == 0 {
			node.Weight = 1
		}
		if word[l-1] == ';' {
			index = 0
			node = nil
		} else {
			index++
		}
	}

	if len(nodes) > 0 {
		for _, n := range nodes {
			instance := s.instanceFactory.General(n.IP, n.Port, n.Weight)
			instances = append(instances, instance)
		}
		s := common.NewService("static_upstream", instances)
		return s, nil
	}
	return nil, ErrorNoInstance

}
