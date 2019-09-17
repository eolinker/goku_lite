package health

import (
	"context"
	"fmt"
	"github.com/eolinker/goku/goku-service/common"
	"net/http"
	"time"
)

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
	url := fmt.Sprintf("http://%s:%d/%s", instance.IP, instance.Port, c.path)
	respone, err := http.Get(url)
	if err != nil {
		return false
	}

	if c.statusCodes[ respone.StatusCode] {
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
				for instanceId, ins := range instances {

					// 处理空列表
					if len(ins) == 0 {
						delete(instances, instanceId)
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
						delete(instances, instanceId)
						continue
					}
					instance := insNew[0]

					if c.check(instance) {
						delete(instances, instanceId)
						for _, in := range insNew {
							in.ChangeStatus(common.InstanceChecking, common.InstanceRun)
						}
					} else {
						count += len(insNew)
						instances[instanceId] = insNew
					}
				}
				c.sum = count
			}
		case instance := <-c.checkChan:
			if instance != nil {
				instances[instance.InstanceId] = append(instances[instance.InstanceId], instance)
				c.sum++
			}
		}
	}
}

func (c *Checker) Check(instance *common.Instance) {
	instance.ChangeStatus(common.InstanceRun, common.InstanceChecking)
	c.checkChan <- instance
}

func (c *Checker) Close() (map[string][]*common.Instance, int) {
	if c.cancelFunc != nil {
		c.cancelFunc()
		c.cancelFunc = nil
	}
	<-c.closeDone
	c.closeDone = nil

	return c.instances, c.sum
}
