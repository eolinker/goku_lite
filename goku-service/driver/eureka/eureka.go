package eureka

import (
	"context"
	"encoding/xml"
	"fmt"
	log "github.com/eolinker/goku/goku-log"
	"github.com/eolinker/goku/goku-service/common"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

type Eureka struct {
	services []*common.Service
	//AppNames  map[string]string
	eurekaUrl       []string
	weightKey       string
	callback        func(services []*common.Service)
	ct              uint64
	cancelFunc      context.CancelFunc
	instanceFactory *common.InstanceFactory
}

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
func (d *Eureka) setConfig(eurekaUrl []string, weightKey string) {
	d.eurekaUrl = eurekaUrl
	d.weightKey = weightKey
}
func (d *Eureka) Driver() string {
	return DriverName
}

func (d *Eureka) SetCallback(callback func(services []*common.Service)) {
	d.callback = callback
}

func (d *Eureka) GetServers() ([]*common.Service, error) {
	return d.services, nil
}

func (d *Eureka) Close() error {
	if d.cancelFunc != nil {
		d.cancelFunc()
		d.cancelFunc = nil
	}
	return nil
}

func (d *Eureka) Open() error {
	d.ScheduleAtFixedRate(time.Second * 5)
	return nil
}

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
			inses = append(inses, d.instanceFactory.General(ins.IpAddr, port, weight))
		}
		server := common.NewService(app.Name, inses)
		services = append(services, server)
	}

	d.callback(services)

}

func (d *Eureka) ScheduleAtFixedRate(second time.Duration) {
	d.run()
	if d.cancelFunc != nil {
		d.cancelFunc()
		d.cancelFunc = nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	d.cancelFunc = cancel
	go d.runTask(second, ctx)
}

func (d *Eureka) runTask(second time.Duration, ctx context.Context) {
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

func (d *Eureka) GetApplications() (*Applications, error) {
	//url := c.eurekaUrl + "/apps"
	url, err := d.getEurekaServerUrl()
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
	var applications *Applications = new(Applications)
	err = xml.Unmarshal(respBody, applications)

	//	log.Info(string(respBody))
	//	log.Info(err, applications)
	return applications, err
}

func (d *Eureka) getEurekaServerUrl() (string, error) {
	ct := atomic.AddUint64(&d.ct, 1)
	size := len(d.eurekaUrl)
	if size == 0 {
		e := NilPointError("eureka url is empty")

		return "", e
	}
	index := int(ct) % size
	url := d.eurekaUrl[index]
	//if strings.LastIndex(url,"/")>-1{
	url = strings.TrimSuffix(url, "/")
	//}
	return url, nil
}

func (d *Eureka) Health() (bool, string) {
	ok, desc := true, "ok"
	i := 0
	for _, u := range d.eurekaUrl {

		url, err := url.Parse(u)
		if err != nil {
			i++
			ok, desc = false, err.Error()
			continue
		}
		healthUrl := url.Scheme + "://" + url.Host + "/health"
		res, err := http.Get(healthUrl)
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
