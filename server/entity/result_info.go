package entity

import (
	"encoding/json"
)

//ResultInfo 结果信息
type ResultInfo struct {
	ResultType string      `json:"type"`
	StatusCode string      `json:"statusCode"`
	ResultKey  string      `json:"resultKey"`
	Result     interface{} `json:"result,omitempty"`
	ResultDesc string      `json:"resultDesc,omitempty"`
}

//String string
func String(info interface{}) string {
	resultInfo, err := json.Marshal(info)
	if err != nil {
		return ""
	}
	return string(resultInfo)
}

//GetResultInfo 获取结果信息
func GetResultInfo(statusCode string, resultType string, resultKey string, resultDesc string, result interface{}, successCount string) map[string]interface{} {
	if resultKey == "" {
		return map[string]interface{}{
			"type":       resultType,
			"statusCode": statusCode,
			"resultDesc": resultDesc,
		}
	}
	if result != nil {
		return map[string]interface{}{
			"type":       resultType,
			"statusCode": statusCode,
			resultKey:    result,
			"resultDesc": resultDesc,
		}
	}
	return map[string]interface{}{
		"type":       resultType,
		"statusCode": statusCode,
		"resultDesc": resultDesc,
	}
}
