package entity

import (
	"encoding/json"
)

type ResultInfo struct {
	ResultType string      `json:"type"`
	StatusCode string      `json:"statusCode"`
	ResultKey  string      `json:"resultKey"`
	Result     interface{} `json:"result,omitempty"`
	ResultDesc string      `json:"resultDesc,omitempty"`
}

func String(info interface{}) string {
	resultInfo, err := json.Marshal(info)
	if err != nil {
		return ""
	}
	return string(resultInfo)
}

func GetResultInfo(statusCode string, resultType string, resultKey string, resultDesc string, result interface{}, successCount string) map[string]interface{} {
	if resultKey == "" {
		return map[string]interface{}{
			"type":       resultType,
			"statusCode": statusCode,
			"resultDesc": resultDesc,
		}
	} else {
		if result != nil {
			return map[string]interface{}{
				"type":       resultType,
				"statusCode": statusCode,
				resultKey:    result,
				"resultDesc": resultDesc,
			}
		} else {
			return map[string]interface{}{
				"type":       resultType,
				"statusCode": statusCode,
				"resultDesc": resultDesc,
			}
		}
	}
}
