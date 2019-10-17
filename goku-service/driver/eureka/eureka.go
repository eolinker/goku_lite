package eureka

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	log "github.com/eolinker/goku-api-gateway/goku-log"
	"github.com/eolinker/goku-api-gateway/goku-service/common"
)

//Eureka eureka
type Eureka struct {
	services []*common.Service
	//AppNames  map[string]string
	eurekaURL       []string
	weightKey       string
	callback        func(services []*common.Service)
	ct              uint64
	cancelFunc      context.CancelFunc
	instanceFactory *common.InstanceFactory
}

//SetConfig setConfig
func (d *Eureka) SetConfig(config string) error {
	tags := strings.Split(config, ";")
	weightKey := ""
	if len(tags) > 1 {
		weightKey = tags[1]
	}

	urls := strings.Split(tags[0], ",")

	d.setConfig(urls, weightKey)

	return nil
}
func (d *Eureka) setConfig(eurekaURL []string, weightKey string) {
	d.eurekaURL = eurekaURL
	d.weightKey = weightKey
}

//Driver driver
func (d *Eureka) Driver() string {
	return DriverName
}

//SetCallback setCallBack
func (d *Eureka) SetCallback(callback func(services []*common.Service)) {
	d.callback = callback
}

//GetServers getServers
func (d *Eureka) GetServers() ([]*common.Service, error) {
	return d.services, nil
}

//Close close
func (d *Eureka) Close() error {
	if d.cancelFunc != nil {
		d.cancelFunc()
		d.cancelFunc = nil
	}
	return nil
}

//Open open
func (d *Eureka) Open() error {
	d.ScheduleAtFixedRate(time.Second * 5)
	return nil
}

//NewEurekaDiscovery 创建Eureka
func NewEurekaDiscovery(config string) *Eureka {
	e := &Eureka{
		services:        nil,
		callback:        nil,
		ct:              0,
		cancelFunc:      nil,
		instanceFactory: common.NewInstanceFactory(),
	}
	e.SetConfig(config)
	return e
}

func (d *Eureka) execCallbacks(apps *Applications) {
	if d.callback == nil {
		return
	}
	if apps == nil {
		d.callback(nil)
		return
	}
	if len(apps.Applications) == 0 {
		d.callback(nil)
		return
	}

	services := make([]*common.Service, 0, len(apps.Applications))
	for _, app := range apps.Applications {
		inses := make([]*common.Instance, 0, len(app.Instances))
		for _, ins := range app.Instances {
			if ins.Status != EurekaStatusUp {
				continue
			}
			weight := 0
			if w, has := ins.Metadata.Map[d.weightKey]; has {
				weight, _ = strconv.Atoi(w)

			}
			if weight == 0 {
				weight = 1
			}
			port := 0
			if ins.Port.Enabled {
				port = ins.Port.Port
			} else if ins.SecurePort.Enabled {
				port = ins.SecurePort.Port
			}
			inses = append(inses, d.instanceFactory.General(ins.IPAddr, port, weight))
		}
		server := common.NewService(app.Name, inses)
		services = append(services, server)
	}

	d.callback(services)

}

//ScheduleAtFixedRate scheduleAtFixedRate
func (d *Eureka) ScheduleAtFixedRate(second time.Duration) {
	d.run()
	if d.cancelFunc != nil {
		d.cancelFunc()
		d.cancelFunc = nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	d.cancelFunc = cancel
	go d.runTask(ctx, second)
}

func (d *Eureka) runTask(ctx context.Context, second time.Duration) {
	timer := time.NewTicker(second)
	for {
		select {
		case <-timer.C:
			d.run()
		case <-ctx.Done():
			return

		}
	}
}

func (d *Eureka) run() {
	apps, err := d.GetApplications()
	if err == nil || apps != nil {
		//d.apps = apps
		d.execCallbacks(apps)
	} else {
		log.Error(err)
	}
}

//GetApplications 获取应用
func (d *Eureka) GetApplications() (*Applications, error) {
	url, err := d.getEurekaServerURL()
	if err != nil {
		return nil, err
	}
	url = fmt.Sprintf("%s/apps", url)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, err
	}
	var applications = new(Applications)
	err = xml.Unmarshal(respBody, applications)

	//	log.Info(string(respBody))
	//	log.Info(err, applications)
	return applications, err
}

func (d *Eureka) getEurekaServerURL() (string, error) {
	ct := atomic.AddUint64(&d.ct, 1)
	size := len(d.eurekaURL)
	if size == 0 {
		e := NilPointError("eureka url is empty")

		return "", e
	}
	index := int(ct) % size
	url := d.eurekaURL[index]
	//if strings.LastIndex(url,"/")>-1{
	url = strings.TrimSuffix(url, "/")
	//}
	return url, nil
}

//Health health
func (d *Eureka) Health() (bool, string) {
	ok, desc := true, "ok"
	i := 0
	for _, u := range d.eurekaURL {

		url, err := url.Parse(u)
		if err != nil {
			i++
			ok, desc = false, err.Error()
			continue
		}
		healthURL := url.Scheme + "://" + url.Host + "/health"
		res, err := http.Get(healthURL)
		if err != nil {
			i++
			ok, desc = false, err.Error()
			continue
		}
		if res == nil || res.StatusCode != http.StatusOK {
			i++
			ok, desc = true, res.Status
			continue
		}

	}

	return ok, desc

}
