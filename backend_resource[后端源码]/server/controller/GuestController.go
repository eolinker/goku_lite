package controller
import (
	"goku-ce-1.0/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"goku-ce-1.0/server/module"
	"regexp"
	"strconv"
)

// 用户登录
func Login(c *gin.Context){
	loginCall := c.PostForm("loginCall")
	loginPassword := c.PostForm("loginPassword")
	if match, _ := regexp.MatchString("^[0-9a-zA-Z][0-9a-zA-Z_]{3,63}$", loginCall);match == false{
		c.JSON(200,gin.H{"type":"guest","statusCode":"120002"})
		return 
	}else if match, _ := regexp.MatchString("^[0-9a-zA-Z]{32}$", loginPassword);match == false{
		c.JSON(200,gin.H{"type":"guest","statusCode":"120004"})
		return 
	}else{
		flag,userID,userToken := module.Login(loginCall,utils.Md5(loginPassword))
		if flag == true{
			userCookie := http.Cookie{Name: "userToken", Value:userToken, Path: "/", MaxAge: 86400}
			userIDCookie := http.Cookie{Name: "userID", Value:strconv.Itoa(userID), Path: "/", MaxAge: 86400}
			// 写入登录信息到redis
			flag = module.ConfirmLoginInfo(userID,loginCall,userToken)
			if flag == true{
				http.SetCookie(c.Writer, &userCookie)
				http.SetCookie(c.Writer, &userIDCookie)
				c.JSON(200,gin.H{"type":"guest","statusCode":"000000",})
				return
			}else{
				c.JSON(200,gin.H{"type":"guest","statusCode":"120000",})
				return 
			}
		}else{
			c.JSON(200,gin.H{"type":"guest","statusCode":"120000",})
			return 
		}
	}
}


