package consul

import (
	"context"
	"time"

	log "github.com/eolinker/goku-api-gateway/goku-log"
	"github.com/eolinker/goku-api-gateway/goku-service/common"
	"github.com/hashicorp/consul/api"
)

//Discovery discovery
type Discovery struct {
	//Config *api.Config

	orgConfig string

	callback func([]*common.Service)
	client   *api.Client
	services []*common.Service

	instanceFactory *common.InstanceFactory
	cancel          context.CancelFunc
}

//SetConfig setConfig
func (d *Discovery) SetConfig(config string) error {
	if d.orgConfig == config {
		return nil
	}
	d.orgConfig = config
	c := api.DefaultConfig()
	c.Address = config
	//d.Config = c
	client, err := api.NewClient(c)
	if err != nil {
		return err
	}
	d.client = client

	return nil

}

//Driver driver
func (d *Discovery) Driver() string {
	return DriverName
}

//SetCallback setCallback
func (d *Discovery) SetCallback(callback func(services []*common.Service)) {
	d.callback = callback
}

//GetServers getServers
func (d *Discovery) GetServers() ([]*common.Service, error) {
	return d.services, nil
}

//Close close
func (d *Discovery) Close() error {
	if d.cancel != nil {
		d.cancel()
	}
	return nil
}

//Open open
func (d *Discovery) Open() error {

	d.ScheduleAtFixedRate(time.Second * 5)
	return nil
}

//NewConsulDiscovery address: [hostName:port]
func NewConsulDiscovery(address string) *Discovery {
	config := api.DefaultConfig()
	config.Address = address

	client, err := api.NewClient(config)
	if err != nil {
		log.Error(err)
		return nil
	}

	cd := &Discovery{
		callback:        nil,
		client:          client,
		services:        nil,
		orgConfig:       address,
		instanceFactory: common.NewInstanceFactory(),
	}

	return cd
}

//GetServicesInTime getServicesInTime
func (d *Discovery) GetServicesInTime() (map[string][]string, map[string][]*api.ServiceEntry, error) {

	q := &api.QueryOptions{}
	services, _, err := d.client.Catalog().Services(q)
	if err != nil {
		return nil, nil, err
	}

	catalogServices := make(map[string][]*api.ServiceEntry)

	for serviceName := range services {
		cs, _, err := d.client.Health().Service(serviceName, "", true, q)
		if err != nil {
			log.Info(err.Error())
			continue
		}
		catalogServices[serviceName] = cs

	}

	return services, catalogServices, nil

}

//ScheduleAtFixedRate scheduleAtFixedRate
func (d *Discovery) ScheduleAtFixedRate(second time.Duration) {
	if d.cancel != nil {
		d.cancel()
		d.cancel = nil
	}
	d.run()
	ctx, cancel := context.WithCancel(context.Background())
	d.cancel = cancel
	go d.runTask(ctx, second)
}

func (d *Discovery) runTask(ctx context.Context, second time.Duration) {
	timer := time.NewTicker(second)
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			{
				return
			}
		case <-timer.C:
			go d.run()
		}
	}
}
func (d *Discovery) run() {
	services, catalogServices, err := d.GetServicesInTime()
	if err == nil || services != nil || catalogServices != nil {
		d.execCallbacks(services, catalogServices)
	} else {
		log.Info(err.Error())
	}
}

func (d *Discovery) execCallbacks(services map[string][]string, catalogServices map[string][]*api.ServiceEntry) {
	if services == nil {
		log.Info("consul services is empty")
		return
	}
	if d.callback == nil {
		return
	}

	serviceList := make([]*common.Service, 0, len(catalogServices))
	for appName, catalogInstances := range catalogServices {

		size := len(catalogInstances)
		if size == 0 {
			continue
		}
		hosts := make([]*common.Instance, size)
		for i, instance := range catalogInstances {
			//h.hostChangedCallback(appName, newHostInstanceByEureka(appName, &instance))
			hosts[i] = d.instanceFactory.General(instance.Node.Address, instance.Service.Port, 1)

		}

		s := common.NewService(appName, hosts)
		serviceList = append(serviceList, s)
	}
	d.services = serviceList
	d.callback(serviceList)

}

//Health health
func (d *Discovery) Health() (bool, string) {
	leader, err := d.client.Status().Leader()
	if err != nil || leader == "" {
		return false, err.Error()
	}

	ok, desc := true, "ok"

	return ok, desc

}
