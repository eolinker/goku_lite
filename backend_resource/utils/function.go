package utils
import (
	"encoding/hex"
	"crypto/sha1"
	"time"
	"strconv"
	"crypto/md5"	
	"goku-ce-1.0/conf"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"strings"
	"os"
	"encoding/json"
	"os/exec"
	"math/rand"
)

// 将string转为int类型
func ConvertString(params string) (int,bool){
	id,err := strconv.Atoi(params)
	if err != nil{
		return 0,false
	}else{
		return id,true
	}
}

func GetRedisConnection() (redis.Conn, error) {
	db, err := strconv.Atoi(conf.Configure["redis_db"])
	if err != nil {
		db = 0
	}
	if _, hasPassword := conf.Configure["redis_password"]; hasPassword {
			return redis.Dial("tcp",
				conf.Configure["redis_host"]+":"+conf.Configure["redis_port"],
				redis.DialPassword(conf.Configure["redis_password"]),
				redis.DialDatabase(db))
	} else {
			return redis.Dial("tcp",
				conf.Configure["redis_host"]+":"+conf.Configure["redis_port"],
				redis.DialDatabase(db))
	}
}



func GetHashKey(first_sail string,args ...string) (string){
	hashKey := ""
	hashKey = hashKey+ strconv.Itoa(int(time.Now().Unix())) + first_sail
	for i:=0;i<len(args);i++{
		hashKey += args[i]
	}
	h := sha1.New()
	h.Write([]byte(hashKey))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5(encodeString string) string{
	h := md5.New()
    h.Write([]byte(encodeString)) 
    return hex.EncodeToString(h.Sum(nil)) // 输出加密结果
}

// 替换数据库名称
func ReplaceDBName(mysql_dbname string) bool{
	data, err := ioutil.ReadFile("./server/conf/eo_gateway.sql")
	if err != nil {
		return false
	}
	
	result := strings.Replace(string(data),"$mysql_dbname",mysql_dbname,-1)
	err = ioutil.WriteFile("./eo_gateway.sql", []byte(result), 0666) //写入文件(字节数组)
	if err != nil{
		return false
	}else{
		return true
	}
}	

// 创建配置文件
func CreateConfigureFile(configureInfo ConfigureInfo) (bool){
	configJson,_ := json.Marshal(configureInfo)
	configString := string(configJson[:])
	err := ioutil.WriteFile("./configure.json", []byte(configString), 0666) //写入文件(字节数组)
	if err != nil{
		return false
	}else{
		return true
	}
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

// 安装数据库
func InstallDB(mysql_username,mysql_password,mysql_host,mysql_port string) (bool){
	sql := "mysql -u" + mysql_username + " -p" + mysql_password + " -h" + mysql_host + " -P" + mysql_port  + " <./eo_gateway.sql"
	cmd := exec.Command("/bin/bash", "-c",sql)
	
	if _, err := cmd.Output(); err != nil {
		return false
	}else{
		return true
	}
}

// 关闭网关服务，重启读取配置文件
func Stop() bool{
	id :=os.Getpid()
	cmd := exec.Command("/bin/bash","-c","kill -HUP " + strconv.Itoa(id))
	if _, err := cmd.Output(); err != nil {
		return false
	}else{
		return true
	}
}

// 启动网关服务
func StartGateway() bool{
	cmd := exec.Command("/bin/bash","-c","go run gateway.go" )
	if _, err := cmd.Output(); err != nil {
		return false
	}else{
		return true
	}
}

// 将[]string转为[]int
func ConvertArray(arr []string) (bool,[]int){
	result := make([]int,0)
	for _,i := range arr{
		res,err := strconv.Atoi(i)
		if err != nil{
			return false,result
		}
		result = append(result,res)
	}
	return true,result
}
