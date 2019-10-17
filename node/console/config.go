package console

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	listener2 "github.com/eolinker/goku-api-gateway/common/listener"
	"github.com/eolinker/goku-api-gateway/config"
)

//ConfigCallbackFunc configCallbackFunc
type ConfigCallbackFunc func(gokuConfig *config.GokuConfig)

var (
	lastConfig *config.GokuConfig
	once       sync.Once

	listener = listener2.New()
)

//GetConfig 获取节点配置
func (c *Console) GetConfig() (*config.GokuConfig, error) {
	if lastConfig != nil {
		return lastConfig, nil
	}

	once.Do(func() {
		listenConfig(c.ctx, c.port, c.adminHost)
	})

	cn := make(chan *config.GokuConfig, 1)
	listener.ListenOnce(func(event interface{}) {
		conf := event.(*config.GokuConfig)
		cn <- conf
	})
	deadline, _ := context.WithTimeout(context.Background(), time.Second*30)

	select {
	case <-deadline.Done():
		return nil, errors.New("get config timeout")
	case conf := <-cn:
		return conf, nil
	}
}

//AddListen 新增监听配置
func (c *Console) AddListen(callback ConfigCallbackFunc) {

	listener.Listen(func(event interface{}) {
		conf := event.(*config.GokuConfig)
		callback(conf)
	})

}

func listenConfig(ctx context.Context, port int, adminHost string) {

	admin := adminHost
	admin = strings.TrimPrefix(admin, "http://")
	admin = strings.TrimSuffix(admin, "/")

	url := fmt.Sprintf("http://%s/version/config/get", admin)

	go func() {

		errNum := 0
		lastVersion := ""
		for {
			select {
			case <-ctx.Done():
				return
			default:
				{

					gokuConfig, err := getConfig(url, port, lastVersion)
					if err != nil {
						errNum++
						time.After(time.Second * time.Duration(errNum))
						continue
					}
					if gokuConfig != nil {
						if lastVersion != gokuConfig.Version {
							lastVersion = gokuConfig.Version
							lastConfig = gokuConfig
							listener.Call(gokuConfig)
						}

					}
				}
			}

		}

	}()

}
func getConfig(url string, port int, lastVersion string) (*config.GokuConfig, error) {
	req, e := http.NewRequest(http.MethodGet, url, nil)

	if e != nil {

		return nil, e
	}

	q := req.URL.Query()
	q.Add("port", strconv.Itoa(port))
	q.Add("version", lastVersion)
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	gConfig := new(config.GokuConfig)

	err = json.Unmarshal(data, gConfig)
	if err != nil {
		return nil, err
	}

	return gConfig, nil
}
