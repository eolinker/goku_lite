package config_log

import "C"
import (
	"encoding/json"
	"strings"

	log "github.com/eolinker/goku-api-gateway/goku-log"
	access_field "github.com/eolinker/goku-api-gateway/server/access-field"
	entity "github.com/eolinker/goku-api-gateway/server/entity/config-log"
)

const (
	//ConsoleLog console日志
	ConsoleLog = "console"
	//NodeLog 节点日志
	NodeLog = "node"
	//AccessLog access日志
	AccessLog = "access"
	//ExpireDefault 默认过期时间
	ExpireDefault = 3
)

var (
	logNames = map[string]int{
		ConsoleLog: 1,
		NodeLog:    1,
		AccessLog:  1,
	}
	//Expires 过期时间选项数组
	Expires = []ValueTitle{
		{
			Value: 3,
			Title: "3天",
		},
		{
			Value: 7,
			Title: "7天",
		},
		{
			Value: 30,
			Title: "30天",
		},
		{
			Value: 90,
			Title: "90天",
		},
		{
			Value: 180,
			Title: "180天",
		},
	}
	//Periods 日志生成周期
	Periods = []NameTitle{
		{
			Name:  log.PeriodDay.String(),
			Title: "天",
		},
		{
			Name:  log.PeriodHour.String(),
			Title: "小时",
		},
	}
	//Levels 日志级别
	Levels = []NameTitle{
		{
			Name:  log.ErrorLevel.String(),
			Title: strings.ToUpper(log.ErrorLevel.String()),
		},
		{
			Name:  log.WarnLevel.String(),
			Title: strings.ToUpper(log.WarnLevel.String()),
		}, {
			Name:  log.InfoLevel.String(),
			Title: strings.ToUpper(log.InfoLevel.String()),
		},
		{
			Name:  log.DebugLevel.String(),
			Title: strings.ToUpper(log.DebugLevel.String()),
		},
	}
)

//NameTitle 名称标题结构体
type NameTitle struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

//ValueTitle 值标题结构体
type ValueTitle struct {
	Value int    `json:"value"`
	Title string `json:"title"`
}

//Param 日志参数
type Param struct {
	Enable bool
	Dir    string
	File   string
	Level  string
	Period string
	Fields string
	Expire int
}

//PutParam put方式所需参数
type PutParam struct {
	Enable bool   `opt:"enable,require"`
	Dir    string `opt:"dir,require"`
	File   string `opt:"file,require"`
	Level  string `opt:"level,require"`
	Period string `opt:"period,require"`
	Expire int    `opt:"expire,require"`
}

//Format 格式化
func (p *PutParam) Format() (*Param, error) {
	l, err := log.ParseLevel(p.Level)
	if err != nil {
		return nil, err
	}
	period, err := log.ParsePeriod(p.Period)
	if err != nil {
		return nil, err
	}
	return &Param{
		Enable: p.Enable,
		Dir:    p.Dir,
		File:   p.File,
		Level:  l.String(),
		Period: period.String(),
		Expire: p.Expire,
		Fields: "",
	}, nil

}

//AccessParam access参数
type AccessParam struct {
	Enable bool   `opt:"enable,require" `
	Dir    string `opt:"dir,require"`
	File   string `opt:"file,require"`
	Period string `opt:"period,require"`
	Fields string `opt:"fields,require"`
	Expire int    `opt:"expire,require"`
}

//Format 格式化
func (p *AccessParam) Format() (*Param, error) {
	period, err := log.ParsePeriod(p.Period)
	if err != nil {
		return nil, err
	}
	return &Param{
		Enable: p.Enable,
		Dir:    p.Dir,
		File:   p.File,
		Level:  "",
		Period: period.String(),
		Fields: p.Fields,
		Expire: p.Expire,
	}, nil
}

//LogConfig 日志配置
type LogConfig struct {
	Name    string       `json:"-"`
	Enable  bool         `json:"enable" opt:"enable" default:"true"`
	Dir     string       `json:"dir" opt:"dir" default:"work/logs/"`
	File    string       `json:"file"`
	Level   string       `json:"level"`
	Levels  []NameTitle  `json:"levels"`
	Period  string       `json:"period"`
	Periods []NameTitle  `json:"periods"`
	Expire  int          `json:"expire"`
	Expires []ValueTitle `json:"expires"`
}

func (c *LogConfig) Read(ent *entity.LogConfig) {

	c.Name = ent.Name
	c.Enable = ent.Enable == 1
	c.Dir = ent.Dir
	c.File = ent.File
	c.Level = ent.Level
	c.Period = ent.Period
	c.Expire = ent.Expire
	if c.Expire < ExpireDefault {
		c.Expire = ExpireDefault
	}
}

//AccessConfig access配置
type AccessConfig struct {
	Name    string         `json:"-"`
	Enable  bool           `json:"enable" opt:"enable" default:"true"`
	Dir     string         `json:"dir" opt:"dir" default:"work/logs/"`
	File    string         `json:"file" opt:"file" default:"access"`
	Period  string         `json:"period"`
	Periods []NameTitle    `json:"periods"`
	Expire  int            `json:"expire"`
	Expires []ValueTitle   `json:"expires"`
	Fields  []*AccessField `json:"fields"`
}

//AccessField access域
type AccessField struct {
	Name   string `json:"name"`
	Select bool   `json:"select"`
	Desc   string `json:"desc"`
}

//InitFields 域初始化
func (c *AccessConfig) InitFields() {
	// 如果有新增的字段，按默认顺序拼接到末尾
	c.Fields = make([]*AccessField, 0, access_field.Size())
	for _, value := range access_field.All() {

		field := &AccessField{
			Name:   value.Key(),
			Select: access_field.IsDefault(value),
			Desc:   value.Info(),
		}
		c.Fields = append(c.Fields, field)
	}
}
func (c *AccessConfig) Read(ent *entity.LogConfig) {

	c.Name = ent.Name
	c.Enable = ent.Enable == 1
	c.Dir = ent.Dir
	c.File = ent.File
	c.Period = ent.Period
	c.Expire = ent.Expire
	fields := make([]*AccessField, 0, access_field.Size())
	e := json.Unmarshal([]byte(ent.Fields), &fields)
	if e != nil {
		c.InitFields()
		return
	}

	all := access_field.CopyKey()

	fieldsTmp := make([]*AccessField, 0, access_field.Size())
	// 筛选有效的字段
	for _, field := range fields {
		if desc, has := all[field.Name]; has {
			field.Desc = desc
			fieldsTmp = append(fieldsTmp, field)
			delete(all, field.Name)
		}
	}

	// 如果有新增的字段，按默认顺序拼接到末尾
	for _, value := range access_field.All() {
		if _, has := all[value.Key()]; has {
			field := &AccessField{
				Name:   value.Key(),
				Select: false,
				Desc:   value.Info(),
			}
			fieldsTmp = append(fieldsTmp, field)
		}
	}

	c.Fields = fieldsTmp

}
