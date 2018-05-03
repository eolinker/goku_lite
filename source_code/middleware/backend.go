package middleware

import (
	"goku-ce/conf"
)

func GetBackendInfo(backendID int,b conf.Backend) (bool,conf.BackendInfo) {
	flag := false
	var backendInfo conf.BackendInfo
	for _,i := range b.Backend {
		if i.BackendID == backendID {
			flag = true
			backendInfo = i
			break
		}
	}
	return flag,backendInfo
}