package utils
import (
	"strings"
	"fmt"
	"encoding/hex"
	"time"
	"crypto/md5"	
    "math/rand"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "os/exec"
    "os"
)

func Md5(encodeString string) string{
	h := md5.New()
    h.Write([]byte(encodeString)) 
    return hex.EncodeToString(h.Sum(nil)) // 输出加密结果
}


//生成随机字符串
func  GetRandomString(num int) string {  
    str := "123456789abcdefghijklmnpqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ"  
    bytes := []byte(str)  
    result := []byte{}  
    r := rand.New(rand.NewSource(time.Now().UnixNano()))  
    for i := 0; i < num; i++ {  
        result = append(result, bytes[r.Intn(len(bytes))])  
    }  
    return string(result)  
}  

func GetVisitCount(url string) map[string]interface{}{
    resp, err := http.Get(url)
    resu := map[string]interface{}{
        "gatewaySuccessCount": 0,
        "gatewayFailureCount": 0,
        "gatewayDayCount": 0,
        "gatewayMinuteCount": 0,
        "lastUpdateTime":time.Now().Format("15:04:05"),
    }
	if err != nil {
        return resu
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
    result := make(map[string]interface{})
	err = json.Unmarshal(body,&result)
    if err != nil {
        return resu
    }
    return result
}

func RestartGatewayService(port string) (bool){
    command := `netstat -tlnp | grep ` + port + `| awk '{print $7}'`
	cmd := exec.Command("/bin/bash", "-c",command)
	
	if out, err := cmd.Output(); err != nil {
        panic(err) 
    }else {
        id  := strings.Split(string(out),"/")[0]
        command = `kill -9 ` + id
        cmd = exec.Command("/bin/bash", "-c",command)
        if _, err := cmd.Output(); err != nil {
            return false
        }
        command = "nohup ./goku-ce -c " + ConfFilepath + "> goku-ce.log 2>&1 &"
        cmd := exec.Command("/bin/bash", "-c",command)
         
        if _, err := cmd.Output(); err != nil {
            return false
        }
        return true
    }
}

func StartGatewayService(port string) (bool){
    command := `netstat -tlnp | grep ` + port + `| awk '{print $7}'`
	cmd := exec.Command("/bin/bash", "-c",command)
	
	if _, err := cmd.Output(); err != nil {
        panic(err)
    }else {
        command = "nohup ./goku-ce -c " + ConfFilepath + "> goku-ce.log 2>&1 &"
        cmd := exec.Command("/bin/bash", "-c",command)
         
        if _, err := cmd.Output(); err != nil {
            return false
        }
        return true
    }
}

// 关闭后端服务
func StopGatewayService(port string,force bool) (bool){
    command := `netstat -tlnp | grep ` + port + `| awk '{print $7}'`
	cmd := exec.Command("/bin/bash", "-c",command)
	
	if out, err := cmd.Output(); err != nil {
        panic(err) 
    }else {
        id  := strings.Split(string(out),"/")[0]
        if force {
            command = `kill -9 ` + id
        } else {
            command = `kill -HUP ` + id
        }
        cmd = exec.Command("/bin/bash", "-c",command)
        if _, err := cmd.Output(); err != nil {
            fmt.Println(port)
            return false
        }
        return true
    }
}

func GetGatewayServiceStatus(port string) (bool){
    command := `netstat -tlnp | grep ` + port + `| awk '{print $7}'`
	cmd := exec.Command("/bin/bash", "-c",command)
	
	if out, err := cmd.Output(); err != nil {
        panic(err) 
    }else {
        id  := strings.Split(string(out),"/")[0]
        if id != "" {
            return true
        } else {
            return false
        }
    }
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
 func CheckFileIsExist(filename string) bool {
	if  _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}else{
		return true
	}
}
