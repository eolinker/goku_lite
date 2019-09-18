package node_common

import "strings"

// 匹配restful参数
func MatchRestful(requestURL string, params []string) string {
	if len(params) == 0 {
		return requestURL
	}
	// 将匹配URI中的query参数清除
	requestURL = InterceptURL(requestURL, "?")
	var postfix bool = false
	// 将URI最后放的"/"去掉
	if string(requestURL[len(requestURL)-1]) == "/" {
		postfix = true
		requestURL = requestURL[:len(requestURL)-1]
	}
	requestArray := strings.Split(requestURL, "/")
	url := "/"
	n := 0
	for _, v := range requestArray {
		if len(v) > 0 {
			if string(v[0]) == ":" {
				url += params[n] + "/"
				n += 1
			} else if string(v[0]) == "{" && string(v[len(v)-1]) == "}" {
				url += params[n] + "/"
				n += 1
			} else {
				url += v + "/"
			}
		}
	}
	if !postfix {
		return url[:len(url)-1]
	} else {
		return url
	}
}

// 匹配URI
func MatchURI(requestPath string, matchURI string) (bool, string, []string) {
	// 将匹配URI中的query参数清除
	matchURI = InterceptURL(matchURI, "?")
	// 将请求参数分组
	if requestPath == matchURI {
		return true, "", nil
	}
	var postfix bool = false
	// 将URI最后的"/"去掉
	if string(requestPath[len(requestPath)-1]) == "/" {
		postfix = true
		requestPath = requestPath[:len(requestPath)-1]
	}
	if string(matchURI[len(matchURI)-1]) == "/" {
		postfix = true
		matchURI = matchURI[:len(matchURI)-1]
	}

	requestArray := strings.Split(requestPath, "/")
	matchArray := strings.Split(matchURI, "/")
	// 比较数组长度
	if len(requestArray) < len(matchArray) {
		return false, "", nil
	}

	n := 0
	param := make([]string, 0)
	for i, v := range matchArray {
		n += 1
		if v == requestArray[i] {
			continue
		} else {
			// 匹配restful参数
			if string(v[0]) == ":" {
				param = append(param, requestArray[i])
				continue
			} else if string(v[0]) == "{" && string(v[len(v)-1]) == "}" {
				param = append(param, requestArray[i])
				continue
			} else {
				return false, "", nil
			}
		}
	}
	splitURL := "/"
	for _, v := range requestArray[n:] {
		splitURL += v + "/"
	}
	if !postfix {
		return true, splitURL[:len(splitURL)-1], param
	} else {
		return true, splitURL, param
	}
}
func InterceptURL(str, substr string) string {
	result := strings.Index(str, substr)
	var rs string
	if result != -1 {
		rs = str[:result]
	} else {
		rs = str
	}
	return rs
}

// 过滤双斜杠
func FilterSlash(uri string) string {
	replaceURI := strings.ReplaceAll(uri, "//", "/")
	return replaceURI
}
