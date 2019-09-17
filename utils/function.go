package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// 将string转为int类型
func ConvertString(params string) (int, bool) {
	id, err := strconv.Atoi(params)
	if err != nil {
		return 0, false
	} else {
		return id, true
	}
}

// 判断ip端口是否合法
func ValidateRemoteAddr(ip string) bool {
	match, err := regexp.MatchString(`^(?:(?:1[0-9][0-9]\.)|(?:2[0-4][0-9]\.)|(?:25[0-5]\.)|(?:[1-9][0-9]\.)|(?:[0-9]\.)){3}(?:(?:1[0-9][0-9])|(?:2[0-4][0-9])|(?:25[0-5])|(?:[1-9][0-9])|(?:[0-9]))\:(([0-9])|([1-9][0-9]{1,3})|([1-6][0-9]{0,4}))$`, ip)
	if err != nil {
		return false
	}
	return match
}

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

func GetHashKey(first_sail string, args ...string) string {
	hashKey := ""
	hashKey = hashKey + strconv.Itoa(int(time.Now().Unix())) + first_sail
	for i := 0; i < len(args); i++ {
		hashKey += args[i]
	}
	h := sha1.New()
	h.Write([]byte(hashKey))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5(encodeString string) string {
	h := md5.New()
	h.Write([]byte(encodeString))
	return hex.EncodeToString(h.Sum(nil)) // 输出加密结果
}

//生成随机字符串
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

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// 关闭网关服务，重启读取配置文件
func Stop() bool {
	id := os.Getpid()
	cmd := exec.Command("/bin/bash", "-c", "kill -HUP "+strconv.Itoa(id))
	if _, err := cmd.Output(); err != nil {
		return false
	} else {
		return true
	}
}

// 启动网关服务
func StartGateway() bool {
	cmd := exec.Command("/bin/bash", "-c", "go run gateway.go")
	if _, err := cmd.Output(); err != nil {
		return false
	} else {
		return true
	}
}

// 将[]string转为[]int
func ConvertArray(arr []string) (bool, []int) {
	result := make([]int, 0)
	for _, i := range arr {
		res, err := strconv.Atoi(i)
		if err != nil {
			return false, result
		}
		result = append(result, res)
	}
	return true, result
}

// 获取MAC地址
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

// 匹配机器码和授权码是否一致
func MatchVerifyCode(verifyCode, mac string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(verifyCode), []byte(mac))
	if err != nil {
		return false
	} else {
		return true
	}
}

// 将机器码加密
func BcryptMAC() string {
	_, mac := GetMac()
	verifyCode, err := bcrypt.GenerateFromPassword([]byte(Md5(mac)), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(verifyCode)
}

// 将数组的值赋给每一个变量
func ConvertArrayToVariable(arr []interface{}, variable ...interface{}) error {
	if len(arr) != len(variable) {
		return errors.New("[ERROR]Fail to convert")
	}
	for i, v := range arr {
		tmp, _ := variable[i].(*int)
		tmpvstr, _ := arr[i].(string)
		tmpv, _ := strconv.Atoi(tmpvstr)
		if v != nil {
			*tmp = tmpv
		} else {
			*tmp = 0
		}
	}
	return nil
}
