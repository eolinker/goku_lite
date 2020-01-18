package dao_version_config

import (
	"encoding/json"
	"strings"

	"github.com/eolinker/goku-api-gateway/config"
)

//GetAPIContent 获取接口信息
func (d *VersionConfigDao)GetAPIContent() ([]*config.APIContent, error) {
	db := d.db
	sql := "SELECT apiID,apiName,IFNULL(protocol,'http'),IFNULL(balanceName,''),IFNULL(targetURL,''),CASE WHEN isFollow = 'true' THEN 'FOLLOW' ELSE targetMethod END targetMethod,responseDataType,requestURL,requestMethod,timeout,alertValve,retryCount,IFNULL(linkApis,''),IFNULL(staticResponse,'') FROM goku_gateway_api"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	apiContents := make([]*config.APIContent, 0, 100)
	defer rows.Close()
	for rows.Next() {
		var apiContent config.APIContent
		var linkApisStr, protocol, balance, targetURL, targetMethod, requestMethod string
		var retryCount int
		linkApis := make([]config.APIStepUIConfig, 0)
		err = rows.Scan(&apiContent.ID, &apiContent.Name, &protocol, &balance, &targetURL, &targetMethod, &apiContent.OutPutEncoder, &apiContent.RequestURL, &requestMethod, &apiContent.TimeOutTotal, &apiContent.AlertThreshold, &retryCount, &linkApisStr, &apiContent.StaticResponse)
		if err != nil {
			return nil, err
		}
		if linkApisStr != "" {
			err = json.Unmarshal([]byte(linkApisStr), &linkApis)
			if err != nil {
				return nil, err
			}
		}

		apiContent.Methods = strings.Split(requestMethod, ",")
		if len(linkApis) < 1 {
			apiContent.Steps = append(apiContent.Steps, &config.APIStepConfig{
				Proto:   protocol,
				Balance: balance,
				Path:    targetURL,
				Method:  targetMethod,
				Encode:  "origin",
				Decode:  apiContent.OutPutEncoder,
				TimeOut: apiContent.TimeOutTotal,
				Retry:   retryCount,
			})
		} else {
			for _, api := range linkApis {
				actions := make([]*config.ActionConfig, 0, 20)
				for _, del := range api.Delete {
					actions = append(actions, &config.ActionConfig{
						ActionType: "delete",
						Original:   del.Origin,
					})
				}
				for _, move := range api.Move {
					actions = append(actions, &config.ActionConfig{
						ActionType: "move",
						Original:   move.Origin,
						Target:     move.Target,
					})
				}
				for _, rename := range api.Rename {
					actions = append(actions, &config.ActionConfig{
						ActionType: "rename",
						Original:   rename.Origin,
						Target:     rename.Target,
					})
				}
				apiContent.Steps = append(apiContent.Steps, &config.APIStepConfig{
					Proto:     api.Proto,
					Balance:   api.Balance,
					Path:      api.Path,
					Body:      api.Body,
					Method:    api.Method,
					Encode:    api.Encode,
					Decode:    api.Decode,
					TimeOut:   api.TimeOut,
					Retry:     api.Retry,
					Group:     api.Group,
					Target:    api.Target,
					WhiteList: api.WhiteList,
					BlackList: api.BlackList,
					Actions:   actions,
				})
			}
		}
		apiContents = append(apiContents, &apiContent)
	}
	return apiContents, nil
}
