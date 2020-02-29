package goku_plugin

import "fmt"

type PluginLogProxy struct {
	pluginName string
}

func GenLogger(pluginName string)Logger  {
	return &PluginLogProxy{
		pluginName:fmt.Sprintf("[plugin:%s]",pluginName),
	}
}

func (l *PluginLogProxy) Debugf(format string, args ...interface{}) {
	format = fmt.Sprintf("%s %s",l.pluginName,format)
	Debugf(format,args...)
}

func (l *PluginLogProxy) Infof(format string, args ...interface{}) {
	format = fmt.Sprintf("%s %s",l.pluginName,format)
	Infof(format,args...)
}

func (l *PluginLogProxy) Warnf(format string, args ...interface{}) {
	format = fmt.Sprintf("%s %s",l.pluginName,format)
	Warnf(format,args...)
}

func (l *PluginLogProxy) Warningf(format string, args ...interface{}) {
	format = fmt.Sprintf("%s %s",l.pluginName,format)
	Warningf(format,args...)
}

func (l *PluginLogProxy) Errorf(format string, args ...interface{}) {
	format = fmt.Sprintf("%s %s",l.pluginName,format)
	Errorf(format,args...)
}

func (l *PluginLogProxy) Debug(args ...interface{}) {
	vs:=make([]interface{},0,len(args)+1)
	vs = append(vs,l.pluginName)
	vs = append(vs,args...)
	Debug(vs...)
}

func (l *PluginLogProxy) Info(args ...interface{}) {
	vs:=make([]interface{},0,len(args)+1)
	vs = append(vs,l.pluginName)
	vs = append(vs,args...)
	Info(vs...)
}

func (l *PluginLogProxy) Warn(args ...interface{}) {
	vs:=make([]interface{},0,len(args)+1)
	vs = append(vs,l.pluginName)
	vs = append(vs,args...)
	Warn(vs...)
}

func (l *PluginLogProxy) Warning(args ...interface{}) {
	vs:=make([]interface{},0,len(args)+1)
	vs = append(vs,l.pluginName)
	vs = append(vs,args...)
	Warning(vs...)
}

func (l *PluginLogProxy) Error(args ...interface{}) {
	vs:=make([]interface{},0,len(args)+1)
	vs = append(vs,l.pluginName)
	vs = append(vs,args...)
	Error(vs...)
}

