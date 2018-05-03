package utils
import (
	"encoding/hex"
	"time"
	"crypto/md5"	
	"math/rand"
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

