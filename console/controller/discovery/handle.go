package discovery

import (
	"net/http"
)

//Handle 处理器
func Handle(prefix string) http.Handler {

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/drivers", getDrivices)

	serveMux.HandleFunc("/add", add)
	serveMux.HandleFunc("/delete", delete)
	serveMux.HandleFunc("/save", edit)
	serveMux.HandleFunc("/info", getInfo)

	serveMux.HandleFunc("/simple", simple)
	serveMux.HandleFunc("/list", list)

	serveMux.HandleFunc("/default", setDefault)

	return http.StripPrefix(prefix, serveMux)

}
