package middleware

import (
    "net/http"
    "strings"
    "goku-ce/goku"
    "goku-ce/request"
    "io/ioutil"
)

func Mapping(res http.ResponseWriter, req *http.Request,param goku.Params,context *goku.Context) {
    // 验证IP是否合法
    f,s := IPLimit(context,res,req) 
    if !f {
        res.WriteHeader(403)
        res.Write([]byte(s))
        return
    }
    f,s = Auth(context,res,req)
    if !f {
        res.WriteHeader(403)
        res.Write([]byte(s))
        return
    }
    f,s = RateLimit(context)
    if !f {
        res.WriteHeader(403)
        res.Write([]byte(s))
        return
    }
    statusCode,body,headers := CreateRequest(context,req,res)
    for key,values := range headers {
        for _,value := range values {
            res.Header().Add(key,value)
        }
    }
    res.WriteHeader(statusCode)
    res.Write(body)
    return
}

// 将请求参数写入请求中
func CreateRequest(g *goku.Context,httpRequest *http.Request,httpResponse http.ResponseWriter) (int,[]byte,map[string][]string) {
    api := g.ApiInfo
    var backendHeaders map[string][]string = make(map[string][]string)
	var backendQueryParams map[string][]string = make(map[string][]string)
    var backendFormParams map[string][]string = make(map[string][]string)
    err := httpRequest.ParseForm()
    if err != nil {
        return 500,[]byte("Parsing Arguments Fail"),make(map[string][]string)
    }
    
    backendMethod := strings.ToUpper(api.ProxyMethod)
    backenDomain := api.BackendPath + api.ProxyURL
    requ,err := request.Method(backendMethod,backenDomain)
    for _, reqParam := range api.ProxyParams {
        var param []string
        isFile := false
		switch reqParam.KeyPosition {
        case "header":
            key := parseHeader(reqParam.Key)
            param = httpRequest.Header[key]
		case "body":
			if httpRequest.Method == "POST" || httpRequest.Method == "PUT" || httpRequest.Method == "PATCH" {
                param = httpRequest.PostForm[reqParam.Key]
                if param == nil {
                    f,fh,err := httpRequest.FormFile(reqParam.Key)
                    if err != nil {
                        panic(err)
                    }
                    defer f.Close()
                    body,err := ioutil.ReadAll(f)
                    if err != nil {
                        panic(err)
                    }
                    requ.AddFile(reqParam.ProxyKey,fh.Filename,body)
                    isFile = true
                }
			} else {
				continue
			}
		case "query":
            param = httpRequest.Form[reqParam.Key]
        }

		if param == nil {
			if reqParam.NotEmpty && !isFile {
				return 400, []byte("Missing required parameters"),make(map[string][]string)
			} else {
				continue
			}
        }
        
		switch reqParam.ProxyKeyPosition {
        case "header":
            key := parseHeader(reqParam.ProxyKey)
			backendHeaders[key] = param
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
			backendHeaders[constParam.Key] = []string{constParam.Value}
		case "body":
			if backendMethod == "POST" || backendMethod == "PUT" || backendMethod == "PATCH" {
				backendFormParams[constParam.Key] = []string{constParam.Value}
			} else {
				backendQueryParams[constParam.Key] = []string{constParam.Value}
            }
        case "query":
			backendQueryParams[constParam.Key] = []string{constParam.Value}
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
    if api.IsRaw {
        body,_ := ioutil.ReadAll(httpRequest.Body)
        requ.SetRawBody([]byte(body))
    } 

    cookies := make(map[string]string)
	for _, cookie := range httpRequest.Cookies() {
		cookies[cookie.Name] = cookie.Value
    }
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

// 修饰请求头
func parseHeader(header string) string {
    headerArray := strings.Split(header,"-")
    result := ""
    for i,h := range headerArray {
        h = strings.Replace(h,"_","",-1)
        result += strings.ToUpper(h[0:1]) + strings.ToLower(h[1:])
        if i + 1 < len(headerArray) {
            result += "-"
        }
    }
    return result
}