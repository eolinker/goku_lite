package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/eolinker/goku-api-gateway/goku-node/common"
)

//GetRandomString 生成随机字符串
func GetRandomString(num int) string {
	str := "123456789abcdefghijklmnpqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//Intercept 获取前缀
func Intercept(str, substr string) string {
	result := strings.Index(str, substr)
	var rs string
	if result != -1 {
		rs = str[:result]
	} else {
		rs = str
	}
	return rs
}

//GetStrateyID 获取策略ID
func GetStrateyID(ctx *common.Context) string {
	if value := ctx.Request().GetHeader("Strategy-Id"); value != "" {
		return value
	}
	if value := ctx.Request().URL().Query().Get("Strategy-Id"); value != "" {
		return value
	}
	if value := ctx.Request().GetForm("Strategy-Id"); value != "" {
		return value
	}

	return ""
}
