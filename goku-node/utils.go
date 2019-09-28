package gokunode

import (
	"math/rand"
	"strings"
	"time"
)

// GetRandomString 生成随机字符串
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

//Intercept 过滤子字符串
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
