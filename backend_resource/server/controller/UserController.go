package controller
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"goku-ce-1.0/server/module"
	"strconv"
	"goku-ce-1.0/utils"
	"regexp"
)

func Logout(c *gin.Context){
	if module.CheckLogin(c) == true{
		userCookie := http.Cookie{Name: "userID", Path: "/", MaxAge: -1}
		http.SetCookie(c.Writer, &userCookie)
		c.JSON(200,gin.H{"type":"guest","statusCode":"000000",})
	}else{
		c.JSON(200,gin.H{"type":"guest","statusCode":"100000",})
	}
}

func GetUserInfo(c *gin.Context){
	if module.CheckLogin(c) == true{
		
		c.JSON(200,gin.H{"type":"guest","statusCode":"000000",})
	}else{
		c.JSON(200,gin.H{"type":"guest","statusCode":"100000",})
	}
}

func EditPassword(c *gin.Context){
	var userID int
	if module.CheckLogin(c) == true{
		result,_ := c.Request.Cookie("userID")
		userID,_ = strconv.Atoi(result.Value)
	}else{
		c.JSON(200,gin.H{"type":"guest","statusCode":"100000",})
		return 
	}
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	if match, _ := regexp.MatchString("^[0-9a-zA-Z]{32}$", oldPassword);match == false{
		c.JSON(200,gin.H{"type":"guest","statusCode":"120006"})
		return 
	}else if match, _ := regexp.MatchString("^[0-9a-zA-Z]{32}$", newPassword);match == false{
		c.JSON(200,gin.H{"type":"guest","statusCode":"120007"})
		return 
	}
	flag := module.EditPassword(userID,utils.Md5(oldPassword),utils.Md5(newPassword))
	if flag == false{
		c.JSON(200,gin.H{"statusCode":"100000","type":"guest",})
		return 
	}else{
		c.JSON(200,gin.H{"type":"guest","statusCode":"000000",})
		return
	}
}

func CheckLogin(c *gin.Context){
	if module.CheckLogin(c) == true{
		c.JSON(200,gin.H{"type":"user","statusCode":"000000",})
	}else{
		c.JSON(200,gin.H{"type":"user","statusCode":"100000",})
	}
}