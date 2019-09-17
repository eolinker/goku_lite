package dao_api

import (
	"fmt"
	"github.com/eolinker/goku/common/database"
	entity "github.com/eolinker/goku/server/entity/node-entity"
)

// 获取API接口列表
func GetAllApi() (map[int]*entity.Api, error) {

	sql := "SELECT `apiID`,`apiName`,`requestMethod`,`requestURL`,`protocol`,`balanceName`,`targetURL`,`targetMethod`,`isFollow`,`stripPrefix`,`stripSlash`,`timeout`,`retryCount`,`alertValve` FROM `goku_gateway_api`   ORDER BY `apiID` asc;"
	stmt, e := database.GetConnection().Prepare(sql)
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	apiList := make(map[int]*entity.Api, 0)
	//获取记录列
	if _, err = rows.Columns(); err != nil {
		return nil, err
	}
	for rows.Next() {
		api := new(entity.Api)
		err = rows.Scan(&api.ApiID, &api.ApiName, &api.RequestMethod, &api.RequestURL, &api.Protocol, &api.BalanceName, &api.TargetURL, &api.TargetMethod, &api.IsFollow, &api.StripPrefix, &api.StripSlash, &api.Timeout, &api.RetryCount, &api.AlertValve)
		if err != nil {
			continue
		}
		if len(api.RequestURL) == 0 || api.RequestURL[0] != '/' {
			api.RequestURL = fmt.Sprint("/", api.RequestURL)
		}
		if len(api.TargetURL) == 0 || api.TargetURL[0] != '/' {
			api.TargetURL = fmt.Sprint("/", api.TargetURL)
		}

		apiList[api.ApiID] = api
	}
	return apiList, nil

}
