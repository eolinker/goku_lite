package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//ValidateRemoteAddr 判断ip端口是否合法
func ValidateRemoteAddr(ip string) bool {
	match, err := regexp.MatchString(`^(?:(?:1[0-9][0-9]\.)|(?:2[0-4][0-9]\.)|(?:25[0-5]\.)|(?:[1-9][0-9]\.)|(?:[0-9]\.)){3}(?:(?:1[0-9][0-9])|(?:2[0-4][0-9])|(?:25[0-5])|(?:[1-9][0-9])|(?:[0-9]))\:(([0-9])|([1-9][0-9]{1,3})|([1-6][0-9]{0,4}))$`, ip)
	if err != nil {
		return false
	}
	return match
}

//InterceptIP 获取IP
func InterceptIP(str, substr string) string {
	result := strings.Index(str, substr)
	var rs string
	if result != -1 {
		rs = str[:result]
	} else {
		rs = str
	}
	return rs
}

//Md5 md5加密
func Md5(encodeString string) string {
	h := md5.New()
	h.Write([]byte(encodeString))
	return hex.EncodeToString(h.Sum(nil)) // 输出加密结果
}

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

//CheckFileIsExist 判断文件是否存在  存在返回 true 不存在返回false
func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

//Stop 关闭网关服务，重启读取配置文件
func Stop() bool {
	id := os.Getpid()
	cmd := exec.Command("/bin/bash", "-c", "kill -HUP "+strconv.Itoa(id))
	if _, err := cmd.Output(); err != nil {
		return false
	}
	return true
}

//GetMac 获取MAC地址
func GetMac() (bool, string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return false, "Poor soul, here is what you got: " + err.Error()
	}
	for _, inter := range interfaces {
		mac := inter.HardwareAddr //获取本机MAC地址
		m := fmt.Sprintf("%s", mac)
		match, err := regexp.MatchString(`[0-9a-f][0-9a-f][:-][0-9a-f][0-9a-f][:-][0-9a-f][0-9a-f][:-][0-9a-f][0-9a-f][:-][0-9a-f][0-9a-f][:-][0-9a-f][0-9a-f]`, m)
		if err != nil {
			return false, ""
		}
		if match {
			return true, string(m)
		}
	}
	return false, ""
}
