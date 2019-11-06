package interpreter

import (
	"fmt"
	"net/http"
	"testing"
)

func TestParse(t *testing.T) {

	tpl := "header:{{header.va}},cookie:{cookie.name},restful:id={{restful.name}},body:{{body.name}}"

	body := map[string]interface{}{
		"name": "bodyName",
	}
	header := http.Header{}
	header.Set("va", "headVA")
	cookie := []*http.Cookie{{Name: "name", Value: "kingsword"}}
	resfult := map[string]string{
		"id":   "1",
		"name": "app",
	}
	variables := NewVariables("{xxxx}", body, header, cookie, resfult, 1)

	interpreter, e := Parse(tpl)
	if e != nil {

		t.Fatal(e)
		return
	}

	path := "/xxx/{name}/:id/:name?a=1"

	interpreterpath, e := ParsePath(path)
	if e != nil {

		t.Fatal(e)
		return
	}
	t.Log("path:\n\t", path)
	t.Log("target:\n\t", interpreterpath.Execution(variables))

	//variables.AppendResponse(header,body,cookie)
	//
	//
	//tpl2:="header:{{header1.va}},cookie:{{cookie1.name}},restful:id={{restful.id}}"
	//interpreter2, e := Parse(tpl2)
	//if e!= nil{
	//
	//	t.Fatal(e)
	//	return
	//}
	//fmt.Println("tpl2:\n\t",tpl2)
	//fmt.Println("target:\n\t",interpreter2.Execution(variables))
}
