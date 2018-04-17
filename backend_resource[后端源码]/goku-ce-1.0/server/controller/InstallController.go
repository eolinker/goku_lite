package controller

import (
    "goku-ce-1.0/server/module"
	"github.com/gin-gonic/gin"
	"goku-ce-1.0/utils"
	"strings"
	"regexp"
)

// 检查数据库是否可以连接
func CheckDBConnect(c *gin.Context) {
	var mysql_host,mysql_port string
	mysql_username := c.PostForm("mysqlUserName")
	mysql_password := c.PostForm("mysqlPassword")
	mysql_string := c.PostForm("mysqlHost")
	mysql_dbname := c.PostForm("mysqlDBName")
	mysql_array := strings.Split(mysql_string,":")
	if len(mysql_array) == 2{
		mysql_host = mysql_array[0]
		mysql_port = mysql_array[1]
	}else if len(mysql_array) == 1{
		mysql_host = mysql_array[0]
		mysql_port = "3306"
	}
	if mysql_host == ""{
		c.JSON(200,gin.H{"statusCode":"200008","type":"install",})
		return
	}

	flag := module.CheckDBConnect(mysql_username,mysql_password,mysql_host,mysql_port,mysql_dbname)
	if flag == true{
		c.JSON(200,gin.H{"statusCode":"000000","type":"install",})
		return 
	}else{
		c.JSON(200,gin.H{"statusCode":"200000","type":"install",})
		return
	}
}

// 检查Redis是否可以连接
func CheckRedisConnect(c *gin.Context) {
	var redis_host,redis_port string
	redis_db := c.PostForm("redisDB")
	redis_password := c.PostForm("redisPassword")
	redis_string := c.PostForm("redisHost")
	redis_array := strings.Split(redis_string,":")
	if len(redis_array) == 2{
		redis_host = redis_array[0]
		redis_port = redis_array[1]
	}else if len(redis_array) == 1{
		redis_host = redis_array[0]
		redis_port = "6379"
	}

	if redis_host == ""{
		c.JSON(200,gin.H{"statusCode":"200009","type":"install",})
		return
	}

	flag := module.CheckRedisConnect(redis_db,redis_password,redis_host,redis_port)
	if flag == true{
		c.JSON(200,gin.H{"statusCode":"000000","type":"install",})
		return 
	}else{
		c.JSON(200,gin.H{"statusCode":"200000","type":"install",})
		return
	}
}

// 创建配置文件
func InstallConfigure(c *gin.Context) {
	var mysql_host,mysql_port,redis_host,redis_port string
	mysql_username := c.PostForm("mysqlUserName")
	mysql_password := c.PostForm("mysqlPassword")
	mysql_string := c.PostForm("mysqlHost")
	mysql_dbname := c.PostForm("mysqlDBName")
	mysql_array := strings.Split(mysql_string,":")
	if len(mysql_array) == 2{
		mysql_host = mysql_array[0]
		mysql_port = mysql_array[1]
	}else if len(mysql_array) == 1{
		mysql_host = mysql_array[0]
		mysql_port = "3306"
	}
	flag := module.CheckDBConnect(mysql_username,mysql_password,mysql_host,mysql_port,mysql_dbname)
	if flag == true{
		redis_db := c.PostForm("redisDB")
		redis_password := c.PostForm("redisPassword")
		redis_string := c.PostForm("redisHost")
		redis_array := strings.Split(redis_string,":")
		if len(redis_array) == 2{
			redis_host = redis_array[0]
			redis_port = redis_array[1]
		}else if len(redis_array) == 1{
			redis_host = redis_array[0]
			redis_port = "6379"
		}
		flag = module.CheckRedisConnect(redis_db,redis_password,redis_host,redis_port)
		if flag == true{
			gatewayPort := c.PostForm("gatewayPort")
			var configureInfo utils.ConfigureInfo
			configureInfo.MysqlUserName = mysql_username
			configureInfo.MysqlPassword = mysql_password
			configureInfo.MysqlHost = mysql_host
			configureInfo.MysqlPort = mysql_port
			configureInfo.MysqlDBName = mysql_dbname
			configureInfo.RedisDB = redis_db
			configureInfo.RedisHost = redis_host
			configureInfo.RedisPort = redis_port
			configureInfo.RedisPassword = redis_password
			configureInfo.GatewayPort = gatewayPort
			configureInfo.IPMinuteVisitLimit = "100"
			configureInfo.DayVisitLimit = "100000"
			configureInfo.DayThroughputLimit = "104857600"
			configureInfo.MinuteVisitLimit = "2000"

			
			flag = utils.CreateConfigureFile(configureInfo)
			if flag == true{
				flag = utils.ReplaceDBName(mysql_dbname)
				if flag{
						// 安装数据库
						flag = utils.InstallDB(mysql_username,mysql_password,mysql_host,mysql_port)
					
					if flag{
						// 数据库安装成功
						c.JSON(200,gin.H{"statusCode":"000000","type":"install",})
						utils.Stop()
						return 
					}else{
						// 数据库安装失败
						c.JSON(200,gin.H{"statusCode":"200005","type":"install",})
						return 
					}	
				}else{
					// 替换数据库名称失败
					c.JSON(200,gin.H{"statusCode":"200004","type":"install",})
					return 
				}
				
			}else{
				// 配置文件无法创建
				c.JSON(200,gin.H{"statusCode":"200003","type":"install",})
				return 
			}
		}else{
			// redis数据库无法连接
			c.JSON(200,gin.H{"statusCode":"200002","type":"install",})
			return
		}
	}else{
		// mysql数据库无法连接
		c.JSON(200,gin.H{"statusCode":"200001","type":"install",})
		return
	}
	
}

func Install(c *gin.Context){
	userName := c.PostForm("userName")
	userPassword := c.PostForm("userPassword")
	if match, _ := regexp.MatchString("^[0-9a-zA-Z][0-9a-zA-Z_]{3,63}$", userName);match == false{
		// 用户名格式非法
		c.JSON(200,gin.H{"type":"guest","statusCode":"120002"})
		return 
	}else if match, _ := regexp.MatchString("^[0-9a-zA-Z]{32}$", userPassword);match == false{
		// 用户密码格式非法
		c.JSON(200,gin.H{"type":"guest","statusCode":"120004"})
		return 
	}else{
		result := module.Register(userName,utils.Md5(userPassword))
		if result == true{
			c.JSON(200,gin.H{"type":"guest","statusCode":"000000"})
			return
		}else{
			// 注册失败
			c.JSON(200,gin.H{"type":"guest","statusCode":"120000"})
			return 
		}
	}
}

func CheckIsInstall(c *gin.Context){
	flag := utils.CheckFileIsExist("configure.json")
	if flag == true{
		c.JSON(200,gin.H{"statusCode":"000000","type":"install",})
		return 
	}else{
		c.JSON(200,gin.H{"statusCode":"200000","type":"install",})
		return
	}
}




