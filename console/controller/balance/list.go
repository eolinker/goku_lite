package balance

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/balance"
	"net/http"
)

//GetBalanceList 获取负载列表
func GetBalanceList(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()

	keyword := r.FormValue("keyword")
	result, err := balance.Search(keyword)
	if err != nil {
		controller.WriteError(w, "260000", "balance", fmt.Sprintf("[ERROR] %s", err.Error()), err)
		return
	}
	controller.WriteResultInfo(w, "balance", "balanceList", result)

	return
}
