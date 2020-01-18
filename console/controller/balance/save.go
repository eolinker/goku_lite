package balance

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/common/auto-form"
	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/balance"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

//SaveBalance 新增负载配置
func SaveBalance(w http.ResponseWriter, r *http.Request) {



	if err := r.ParseForm(); err != nil {
		controller.WriteError(w, "260000", "data", "[param_check] Parse form body error | 解析form表单参数错误", err)
		return
	}
	param := new(balance.Param)
	err := auto.SetValues(r.PostForm, param)
	if err != nil {
		controller.WriteError(w, "260000", "data", fmt.Sprintf("[param_check] %s", err.Error()), err)
		return
	}

	restlt, err := balance.Save(param)

	if err != nil {
		mysqlError, ok := err.(*mysql.MySQLError)
		if ok && mysqlError.Number == 1062 {
			controller.WriteError(w, "260002", "balance", "负载名重复", err)
			return
		}
		controller.WriteError(w, "260000", "balance", restlt, err)
		return
	}

	controller.WriteResultInfo(w, "balance", "", nil)
	return
}
