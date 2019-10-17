package health

import (
	"strconv"
	"strings"
	"time"

	"github.com/eolinker/goku-api-gateway/goku-service/common"
)

//CheckHandler checkHandler
type CheckHandler interface {
	Open(path string, statusCodes string, second int, timeout time.Duration)
	Check(instance *common.Instance)
	IsNeedCheck() bool
	Close() []*common.Instance
}

//CheckBox checkBox
type CheckBox struct {
	isNeedCheck bool
	statusCodes map[string]bool
	checker     *Checker
}

//Open open
func (c *CheckBox) Open(path string, statusCodes string, second int, timeout time.Duration) {

	old := c.checker

	checker := new(Checker)

	checker.path = strings.TrimPrefix(path, "/")
	status := make(map[int]bool)
	for _, s := range strings.Split(statusCodes, ",") {
		code, e := strconv.Atoi(s)
		if e != nil {
			status[code] = true
		}

	}
	if len(status) == 0 {
		status[200] = true
	}
	checker.statusCodes = status
	checker.second = second
	if checker.second < 5 {
		checker.second = 5
	}
	checker.timeout = timeout

	if checker.timeout < time.Millisecond*100 {
		checker.timeout = time.Millisecond * 100
	}

	if old != nil {
		sources, _ := old.Close()
		checker.instances = sources
	}

	checker.Open()
	c.checker = checker
	c.isNeedCheck = true
}

//Check check
func (c *CheckBox) Check(instance *common.Instance) {
	if !c.isNeedCheck {
		return
	}
	if c.checker != nil {
		c.checker.Check(instance)
	}
}

//IsNeedCheck isNeedCheck
func (c *CheckBox) IsNeedCheck() bool {

	return c.isNeedCheck && c.checker != nil
}

//Close close
func (c *CheckBox) Close() []*common.Instance {
	if c.checker == nil {
		return nil
	}
	c.isNeedCheck = false
	oldInstances, count := c.checker.Close()
	c.checker = nil
	if count > 0 {
		instances := make([]*common.Instance, 0, count)

		for _, ins := range oldInstances {

			if len(ins) > 0 {
				instances = append(instances, ins...)
			}
		}

		return instances
	}

	return nil
}
