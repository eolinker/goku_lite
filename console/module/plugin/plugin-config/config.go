package plugin_config

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func isEnd(r rune) bool {
	if r == ',' || r == '\n' || r == '{' || r == '}' {
		return true
	}
	return false
}

var allConfigOfPlugin map[string]interface{}

//CheckConfig 检查插件配置是否有效
func CheckConfig(pluginName string, config []byte) (bool, error) {
	v, has := allConfigOfPlugin[pluginName]
	if has {
		err := json.Unmarshal(config, v)
		if err != nil {
			switch v := err.(type) {
			case *json.SyntaxError:
				{
					end := int64(bytes.IndexFunc(config[v.Offset:], isEnd))
					if end == -1 {
						end = int64(len(config) - 1)
					} else {
						end = end + v.Offset
					}
					start := 0
					if v.Offset > 0 {
						start = bytes.LastIndexFunc(config[:v.Offset], isEnd)
					}
					if start == -1 {
						start = 0
					}

					return false, fmt.Errorf("json格式错误：%s", string(config[start:end]))
				}
			case *json.UnmarshalTypeError:
				{
					return false, fmt.Errorf("数据类型不正确:\"%s\":%s", v.Field, v.Value)
				}
			}

			return false, err
		}
	}
	return true, nil
}
