package middleware

import (
    "net/http"
    "strings"
    "goku-ce/goku"
    "goku-ce/conf"
    "goku-ce/request"
)

func Mapping(g *goku.Goku,res http.ResponseWriter, req *http.Request) (bool,string){
    url := req.RequestURI
    requestURI := strings.Split(url,"/")
    if len(requestURI) == 2 {
        if requestURI[1] == "" {
            res.WriteHeader(404)
            res.Write([]byte("Lack gatewayAlias"))
            } else {
            return false,"Lack gatewayAlias"
                res.WriteHeader(404)
                return false,"Lack StrategyKey"
            }
        }
        gatewayAlias := requestURI[1]
        strategyKey := requestURI[2]
        urlLen := len(gatewayAlias) + len(strategyKey) + 2
        flag := false
        for _,m := range g.ServiceConfig.GatewayList{
            if m.GatewayAlias == gatewayAlias{
                for _,i := range m.StrategyList.Strategy{
                    if i.StrategyID == strategyKey{
                    flag = true
                    f,r := IPLimit(m,i,res,req)
                    if !f {
                        res.Write([]byte(r))
                        return false,r
                    }

                    f,r = Auth(i,res,req)
                    if !f {
                        res.Write([]byte(r))
                        return false,r
                    } 

                    f,r = RateLimit(g,i)
                    if !f {
                        res.Write([]byte(r))
                        return false,r
                    }                    
                    break
                }
			}
		}
		if flag {
			for _,i := range m.ApiList.Apis{
                if i.RequestURL == url[urlLen:]{
                    // 验证请求
                    if !validateRequest(i,req){
                        res.WriteHeader(404)
                        res.Write([]byte("Error Request Method!"))
                        return false,"Error Request Method!"
                    }
                    
                    // 验证后端信息是否存在
                    f,r := GetBackendInfo(i.BackendID,m.BackendList)
                    if !f {
                        res.WriteHeader(404)
                        res.Write([]byte("Backend config are not exist!"))
                        return false,"Backend config are not exist!"
                    }
                    

                    _,response,httpResponseHeader := CreateRequest(i,r,req,res)
                    for key, values := range httpResponseHeader {
                        for _, value := range values {
                            res.Header().Add(key,value)
                        }
                    }
                    res.Write(response)
                    return true,string(response)
                }
            }
		}
	}
    res.Write([]byte("URI Not Found"))
    return false,"URI Not Found"
}

// 验证协议及请求参数
func validateRequest(api conf.ApiInfo, req *http.Request) bool{
    flag := false
    for _,method := range api.RequestMethod{
        if !(strings.ToUpper(method) == req.Method){
            flag = true
            break
        }
    }
    return flag
}

// 将请求参数写入请求中
func CreateRequest(api conf.ApiInfo,i conf.BackendInfo,httpRequest *http.Request,httpResponse http.ResponseWriter) (int,[]byte,map[string][]string) {
    var backendHeaders map[string][]string = make(map[string][]string)
	var backendQueryParams map[string][]string = make(map[string][]string)
    var backendFormParams map[string][]string = make(map[string][]string)
    err := httpRequest.ParseForm()
    if err != nil {
        return 500,[]byte("Fail to Parse Args"),make(map[string][]string)
    }
    
    backendMethod := api.ProxyMethod
    backenDomain := i.BackendPath + api.ProxyURL
    requ,err := request.Method(strings.ToUpper(backendMethod),backenDomain)
    for _, reqParam := range api.ProxyParams {
		var param []string
		switch reqParam.KeyPosition {
		case "header":
			param = httpRequest.Header[reqParam.Key]
		case "body":
			if httpRequest.Method == "POST" || httpRequest.Method == "PUT" || httpRequest.Method == "PATCH" {
                param = httpRequest.PostForm[reqParam.Key]
			} else {
				continue
			}
		case "query":
			param = httpRequest.Form[reqParam.Key]
        }
		if param == nil {
			if reqParam.NotEmpty {
				return 400, []byte("missing required parameters"),make(map[string][]string)
			} else {
				continue
			}
        }
        
		switch reqParam.ProxyKeyPosition {
		case "header":
			backendHeaders[reqParam.ProxyKey] = param
        case "body":
			if backendMethod == "POST" || backendMethod == "PUT" || backendMethod == "PATCH" {
				backendFormParams[reqParam.ProxyKey] = param
			}
		case "query":
			backendQueryParams[reqParam.ProxyKey] = param
		}
    }
    
    for _, constParam := range api.ConstantParams {
		switch constParam.Position {
		case "header":
			backendHeaders[constParam.Key] = []string{constParam.Key}
		case "body":
			
			if backendMethod == "POST" || backendMethod == "PUT" || backendMethod == "PATCH" {
				backendFormParams[constParam.Key] = []string{constParam.Value}
			} else {
				backendQueryParams[constParam.Key] = []string{constParam.Value}
			}
		}
    }
    
    if err != nil{
        panic(err)
    }

    for key, values := range backendHeaders {
		requ.SetHeader(key, values...)
	}
	for key, values := range backendQueryParams {
		requ.SetQueryParam(key, values...)
	}
	for key, values := range backendFormParams {
		requ.SetFormParam(key, values...)
    }
    if api.ProxyBodyType == "raw" {
        requ.SetRawBody([]byte(api.ProxyBody))
    } else if api.ProxyBodyType == "json" {
        requ.SetJSON(api.ProxyBody)
    }

    
    cookies := make(map[string]string)
	for _, cookie := range httpRequest.Cookies() {
		cookies[cookie.Name] = cookie.Value
    }
    // requ.SetHeader("Cookie",cookies)

    res, err := requ.Send()
    if err != nil {
        return 500,[]byte(""),make(map[string][]string)
    }

    httpResponseHeader := httpResponse.Header()
	for key, _ := range httpResponseHeader {
		httpResponseHeader[key] = nil
    }
	for key, values := range res.Headers() {
		httpResponseHeader[key] = values
    }

    return res.StatusCode(), res.Body(),httpResponseHeader
} 
