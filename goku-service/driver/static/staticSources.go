package static

import (
	"errors"
	"fmt"
	"github.com/eolinker/goku-api-gateway/goku-service/discovery"
	"github.com/eolinker/goku-api-gateway/goku-service/health"
	"time"

	"github.com/eolinker/goku-api-gateway/goku-service/common"
	"strconv"
	"strings"
	"unicode"
)

var ErrorNoInstance = errors.New("no instance")

type Sources struct {
	name string

	discovery          *StaticDiscovery
	healthCheckHandler health.CheckHandler
	instanceFactory    *common.InstanceFactory
}

func (s *Sources) SetHealthConfig(conf *discovery.HealthCheckConfig) {
	if conf == nil || !conf.IsHealthCheck {
		s.Close()
		return
	}

	s.healthCheckHandler.Open(
		conf.Url,
		conf.StatusCode,
		conf.Second,
		time.Duration(conf.TimeOutMill)*time.Millisecond)
}

func (s *Sources) Close() {
	instances := s.healthCheckHandler.Close()
	for _, instance := range instances {
		instance.ChangeStatus(common.InstanceChecking, common.InstanceRun)
	}
}

func (s *Sources) CheckDriver(driverName string) bool {

	return driverName == DriverName
}

func (s *Sources) SetDriverConfig(config string) error {
	return nil
}

func (s *Sources) GetApp(app string) (*common.Service, health.CheckHandler, bool) {
	service, e := s.decode(app)
	if e != nil {
		return nil, nil, false
	}
	return service, s.healthCheckHandler, true
}

func NewStaticSources(name string) *Sources {
	return &Sources{

		name:               name,
		discovery:          new(StaticDiscovery),
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
	var node *Node = nil
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
				if node.Port == 0 {
					node.Port = 80
				}
				//node.InstanceId = fmt.Sprintf("%s:%d",node.IP,node.Port)
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
