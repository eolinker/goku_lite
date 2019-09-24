package consul

import (
	"context"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"github.com/eolinker/goku-api-gateway/goku-service/common"
	"github.com/hashicorp/consul/api"
	"time"
)

type ConsulDiscovery struct {
	//Config *api.Config

	orgConfig string

	callback func([]*common.Service)
	client   *api.Client
	services []*common.Service

	instanceFactory *common.InstanceFactory
	cancel          context.CancelFunc
}

func (d *ConsulDiscovery) SetConfig(config string) error {
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

func (d *ConsulDiscovery) Driver() string {
	return DriverName
}

func (d *ConsulDiscovery) SetCallback(callback func(services []*common.Service)) {
	d.callback = callback
}

func (d *ConsulDiscovery) GetServers() ([]*common.Service, error) {
	return d.services, nil
}

func (d *ConsulDiscovery) Close() error {
	if d.cancel != nil {
		d.cancel()
	}
	return nil
}

func (d *ConsulDiscovery) Open() error {

	d.ScheduleAtFixedRate(time.Second * 5)
	return nil
}

//address: [hostName:port]
func NewConsulDiscovery(address string) *ConsulDiscovery {
	config := api.DefaultConfig()
	config.Address = address

	client, err := api.NewClient(config)
	if err != nil {
		log.Error(err)
		return nil
	}

	cd := &ConsulDiscovery{
		callback:        nil,
		client:          client,
		services:        nil,
		orgConfig:       address,
		instanceFactory: common.NewInstanceFactory(),
	}

	return cd
}

func (d *ConsulDiscovery) GetServicesInTime() (map[string][]string, map[string][]*api.ServiceEntry, error) {

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

func (d *ConsulDiscovery) ScheduleAtFixedRate(second time.Duration) {
	if d.cancel != nil {
		d.cancel()
		d.cancel = nil
	}
	d.run()
	ctx, cancel := context.WithCancel(context.Background())
	d.cancel = cancel
	go d.runTask(second, ctx)
}

func (d *ConsulDiscovery) runTask(second time.Duration, ctx context.Context) {
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
func (d *ConsulDiscovery) run() {
	services, catalogServices, err := d.GetServicesInTime()
	if err == nil || services != nil || catalogServices != nil {
		d.execCallbacks(services, catalogServices)
	} else {
		log.Info(err.Error())
	}
}

func (d *ConsulDiscovery) execCallbacks(services map[string][]string, catalogServices map[string][]*api.ServiceEntry) {
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

func (d *ConsulDiscovery) Health() (bool, string) {
	leader, err := d.client.Status().Leader()
	if err != nil || leader == "" {
		return false, err.Error()
	}

	ok, desc := true, "ok"

	return ok, desc

}
