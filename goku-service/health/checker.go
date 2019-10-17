package health

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/eolinker/goku-api-gateway/goku-service/common"
)

//Checker checker
type Checker struct {
	path    string
	second  int
	timeout time.Duration

	instances  map[string][]*common.Instance
	sum        int
	cancelFunc context.CancelFunc

	statusCodes map[int]bool
	closeDone   chan int
	checkChan   chan *common.Instance
}

//Open open
func (c *Checker) Open() {
	if c.cancelFunc != nil {
		return
	}

	if c.checkChan == nil {
		c.checkChan = make(chan *common.Instance, 2)
	}

	ctx, cancel := context.WithCancel(context.Background())
	c.cancelFunc = cancel
	c.closeDone = make(chan int)
	go c.doloop(ctx, c.closeDone)
}
func (c *Checker) check(instance *common.Instance) bool {
	// todo 这里有个问题，没有指定协议，只能通过端口进行简单判定，如果没有设置端口或者设置自定义的https的端口，会导致无法识别到https
	server := instance.IP
	if instance.Port != 0 {
		server = fmt.Sprintf("%s:%d", instance.IP, instance.Port)
	}
	protocol := "http"
	if instance.Port == 443 {
		protocol = "https"
	}

	url := fmt.Sprintf("%s://%s/%s", protocol, server, c.path)
	response, err := http.Get(url)
	if err != nil {
		return false
	}

	if c.statusCodes[response.StatusCode] {
		return true
	}
	return false
}
func (c *Checker) doloop(ctx context.Context, closeDone chan int) {
	defer close(closeDone)

	t := time.NewTicker(time.Duration(c.second) * time.Second)

	defer t.Stop()

	instances := c.instances
	if instances == nil {
		instances = make(map[string][]*common.Instance)
	}

	for {
		select {
		case <-ctx.Done():
			c.instances = instances

			return
		case <-t.C:
			{
				count := 0
				for instanceID, ins := range instances {

					// 处理空列表
					if len(ins) == 0 {
						delete(instances, instanceID)
						continue
					}

					// 筛选需要检查的实例
					insNew := make([]*common.Instance, 0, len(ins))
					for _, instance := range ins {
						if instance.CheckStatus(common.InstanceChecking) {
							insNew = append(insNew, instance)
						}
					}

					// 移除没有需要待检查的实例id
					if len(insNew) == 0 {
						delete(instances, instanceID)
						continue
					}
					instance := insNew[0]

					if c.check(instance) {
						delete(instances, instanceID)
						for _, in := range insNew {
							in.ChangeStatus(common.InstanceChecking, common.InstanceRun)
						}
					} else {
						count += len(insNew)
						instances[instanceID] = insNew
					}
				}
				c.sum = count
			}
		case instance := <-c.checkChan:
			if instance != nil {
				instances[instance.InstanceID] = append(instances[instance.InstanceID], instance)
				c.sum++
			}
		}
	}
}

//Check check
func (c *Checker) Check(instance *common.Instance) {
	instance.ChangeStatus(common.InstanceRun, common.InstanceChecking)
	c.checkChan <- instance
}

//Close close
func (c *Checker) Close() (map[string][]*common.Instance, int) {
	if c.cancelFunc != nil {
		c.cancelFunc()
		c.cancelFunc = nil
	}
	<-c.closeDone
	c.closeDone = nil

	return c.instances, c.sum
}
