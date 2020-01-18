package dao_version_config

import (
	"github.com/eolinker/goku-api-gateway/config"
)

//GetGatewayBasicConfig GetGatewayBasicConfig
func (d *VersionConfigDao) GetGatewayBasicConfig() (*config.Gateway, error) {
	db := d.db
	sql := "SELECT skipCertificate FROM goku_gateway;"

	var g config.Gateway
	err := db.QueryRow(sql).Scan(&g.SkipCertificate)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
