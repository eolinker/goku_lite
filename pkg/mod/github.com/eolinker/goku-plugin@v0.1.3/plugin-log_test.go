package goku_plugin

import "testing"
const PluginNameTest = "goku_test"
func TestGenLogger(t *testing.T) {
	l:=GenLogger(PluginNameTest)
	l.Debug("Debug")
	l.Info("info")
	l.Warn("warn")
	l.Warning("warning")
	l.Error("error")

	l.Debugf("Debugf")
	l.Infof("Infof")
	l.Warnf("Warnf")
	l.Warningf("warningf")
	l.Errorf("errorf")
}