package config_log

import (
	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao"
)

var (
	configLogDao dao.ConfigLogDao
)

func init() {
	pdao.Need(&configLogDao)
}