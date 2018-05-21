package dao

import (
	"time"
	"goku-ce/server/conf"
	"os"
	"gopkg.in/yaml.v2"
	"sort"
)

// 新增网关
func AddGateway(gatewayName,gatewayAlias string) (bool) {
	gateway := conf.ParseGatewayInfo(conf.GlobalConf.GatewayConfPath)
	_,ok := gateway[gatewayAlias]
	if ok {
		return false
	} else {

		pthSep := string(os.PathSeparator)
		gatewayDir := conf.GlobalConf.GatewayConfPath + pthSep + gatewayAlias

		err := os.Mkdir(gatewayDir, os.ModePerm)  
		if err != nil {
			panic(err)
		}
		apiFile, err := os.Create(gatewayDir + pthSep + "api.conf");
		if err != nil {
			panic(err);
		}
		apiFile.Close();

		apiGroupFile, err := os.Create(gatewayDir + pthSep + "api_group.conf");
		if err != nil {
			panic(err);
		}
		apiGroupFile.Close();

		strategyFile, err := os.Create(gatewayDir + pthSep + "strategy.conf");
		if err != nil {
			panic(err);
		}
		strategyFile.Close();

		backendFile, err := os.Create(gatewayDir + pthSep + "backend.conf");
		if err != nil {
			panic(err);
		}
		backendFile.Close();
		now := time.Now().Format("2006-01-02 15:04:05")
		gateway[gatewayAlias] = &conf.GatewayInfo{
			GatewayName : gatewayName,
			GatewayAlias : gatewayAlias,
			GatewayStatus : "on",
			ApiConfPath : gatewayDir + pthSep + "api.conf",
			ApiGroupConfPath : gatewayDir + pthSep + "api_group.conf",
			StrategyConfPath : gatewayDir + pthSep + "strategy.conf",
			BackendConfPath : gatewayDir + pthSep + "backend.conf",
			UpdateTime : now,
			CreateTime : now,
		}
		gatewayConf,err := yaml.Marshal(gateway[gatewayAlias])
		if err != nil {
			panic(err)
		}
		conf.WriteConfigToFile(gatewayDir + pthSep + "gateway.conf",gatewayConf)
		
		return true
	}
}

// 修改网关信息
func EditGateway(gatewayName,gatewayAlias,oldGatewayAlias string) bool {
	gateway := conf.ParseGatewayInfo(conf.GlobalConf.GatewayConfPath)
	_,ok := gateway[oldGatewayAlias]
	if !ok {
		return false
	} else {
		pthSep := string(os.PathSeparator)
		gatewayDir := conf.GlobalConf.GatewayConfPath + pthSep
		if oldGatewayAlias != gatewayAlias {
			// 如果网关别名修改，重命名文件夹
			err := os.Rename(gatewayDir + oldGatewayAlias, gatewayDir + gatewayAlias)
			if err != nil {
				panic(err)
			}
			gateway[gatewayAlias] = gateway[oldGatewayAlias]
			gateway[gatewayAlias].GatewayName = gatewayName
			gateway[gatewayAlias].GatewayAlias = gatewayAlias
			gateway[gatewayAlias].ApiConfPath = gatewayDir + gatewayAlias + pthSep + "api.conf"
			gateway[gatewayAlias].ApiGroupConfPath = gatewayDir + gatewayAlias + pthSep + "api_group.conf"
			gateway[gatewayAlias].StrategyConfPath = gatewayDir + gatewayAlias + pthSep + "strategy.conf"
			gateway[gatewayAlias].BackendConfPath = gatewayDir + gatewayAlias + pthSep + "backend.conf"
			delete(gateway,oldGatewayAlias)
		} else {
			gateway[oldGatewayAlias].GatewayName = gatewayName
		}
		gatewayConf,_ := yaml.Marshal(gateway[gatewayAlias])
		conf.WriteConfigToFile(gatewayDir + gatewayAlias + pthSep + "gateway.conf",gatewayConf)
		return true
	}
}

// 删除网关
func DeleteGateway(gatewayAlias string) bool {
	gateway := conf.ParseGatewayInfo(conf.GlobalConf.GatewayConfPath)
	_,ok := gateway[gatewayAlias]
	if !ok {
		return false
	} else {
		pthSep := string(os.PathSeparator)
		gatewayDir := conf.GlobalConf.GatewayConfPath + pthSep + gatewayAlias
		err := os.RemoveAll(gatewayDir)
		if err != nil {
			panic(err)
		}
		return true
	}
}


// 获取网关列表
func GetGatewayList() (bool,[]map[string]interface{}) {
	gateway := conf.ParseGatewayInfo(conf.GlobalConf.GatewayConfPath)
	gatewayList := make([]map[string]interface{},0)

	for _,value := range gateway {
		gatewayInfo := map[string]interface{}{
			"gatewayName" : value.GatewayName,
			"gatewayAlias": value.GatewayAlias,
			"gatewayStatus": value.GatewayStatus,
			"updateTime": value.UpdateTime,
		}
		gatewayList = append(gatewayList,gatewayInfo)
	}
	sort.Sort(sort.Reverse(conf.GatewaySlice(gatewayList)))
	return true,gatewayList
}


// 获取网关信息
func GetGatewayInfo(gatewayAlias string) (bool,map[string]interface{}) {
	gateway := conf.ParseGatewayInfo(conf.GlobalConf.GatewayConfPath)
	value,ok := gateway[gatewayAlias]
	if !ok {
		return false,make(map[string]interface{})
	}
	gatewayInfo := map[string]interface{}{
		"gatewayName" : value.GatewayName,
		"gatewayAlias": value.GatewayAlias,
		"gatewayStatus": value.GatewayStatus,
		"updateTime": value.UpdateTime,
		"gatewayProto": "HTTP",
		"gatewayPort": conf.GlobalConf.Port,
		"gatewayHost": conf.GlobalConf.Host,
	}
	return true,gatewayInfo
}

// 获取网关配置路径
func GetGatewayConfPath(gatewayAlias string) (bool,map[string]string) {
	gateway := conf.ParseGatewayInfo(conf.GlobalConf.GatewayConfPath)
	value,ok := gateway[gatewayAlias]
	if !ok {
		return false,make(map[string]string)
	}
	return true, map[string]string{
		"apiConfPath" : value.ApiConfPath,
		"apiGroupConfPath" : value.ApiGroupConfPath,
		"backendConfPath" : value.BackendConfPath,
		"strategyConfPath" : value.StrategyConfPath,
	}
}

// 检查网关别名是否存在
func CheckGatewayAliasIsExist(gatewayAlias string) bool {
	gateway := conf.ParseGatewayInfo(conf.GlobalConf.GatewayConfPath)
	_,ok := gateway[gatewayAlias]
	if !ok {
		return false
	} else {
		return true
	}
}

