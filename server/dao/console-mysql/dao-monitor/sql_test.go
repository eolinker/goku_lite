package dao_monitor

import (
	"fmt"
	monitor_key "github.com/eolinker/goku/server/monitor/monitor-key"
	"testing"
)

func TestGenSql(t *testing.T)  {
	keys := monitor_key.Keys()
	fields := make([]string,len(keys))

	for i:=range keys{
		fields[i] = fmt.Sprintf("`%s`",keys[i].String())
	}
	fmt.Println("# save ")
	fmt.Println(genSqlSave(fields))

	fmt.Println("#全局")
	fmt.Println(genSqlSelectGateway(fields))

	fmt.Println("#全局 按小时")
	fmt.Println(genSqlSelectGatewayByHour(fields))

	fmt.Println("#API")
	fmt.Println(genSqlSelectAPI(fields))

	fmt.Println("#API 下 策略")
	fmt.Println(genSqlSelectStrategyOfApi(fields))

	fmt.Println("#策略")
	fmt.Println(genSqlSelectStrategy(fields))

	fmt.Println("#策略内API")
	fmt.Println(genSqlSelectApiOfStrategy(fields))
}