package utils


import (
	"encoding/json"
)

type ResultInfo struct {
	ResultType				string					`json:"type"`
	StatusCode				string					`json:"statusCode"`
	ResultKey				string					`json:"resultKey"`
	Result					interface{} 			`json:"result,omitempty"`
}

func String(info interface{}) string {
	resultInfo, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	return string(resultInfo)
}

func GetResultInfo(statusCode string, resultType string,resultKey string, result interface{}, successCount string) map[string]interface{} {
	if resultKey == "" {
		return map[string]interface{}{
			"type": resultType,
			"statusCode":statusCode,
		}
	} else {
		if result != nil {
			return map[string]interface{}{
				"type": resultType,
				"statusCode":statusCode,
				resultKey: result,
			}
		} else {
			return map[string]interface{}{
				"type": resultType,
				"statusCode":statusCode,
			}
		}
	}
}
