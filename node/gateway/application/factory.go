package application

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/eolinker/goku-api-gateway/config"
)

var (
	//ErrorInvalidAPI 无效API错误
	ErrorInvalidAPI = errors.New("invalid api")
)

//Factory factory
type Factory struct {
	apiContents map[int]*config.APIContent

	cache map[string]Application
}

//NewFactory create new factory
func NewFactory(apis map[int]*config.APIContent) *Factory {
	return &Factory{
		apiContents: apis,
		cache:       make(map[string]Application),
	}
}

//GenApplication 通过配置生成应用
func (f *Factory) GenApplication(cfg *config.APIOfStrategy) (Application, error) {

	apiContent, has := f.apiContents[cfg.ID]
	if !has {
		return nil, ErrorInvalidAPI
	}
	switch len(apiContent.Steps) {
	case 0:
		{
			key := fmt.Sprintf("Empty:%d", cfg.ID)
			app, has := f.cache[key]
			if !has {
				app = NewEmptyApplication(apiContent.StaticResponse)
				f.cache[key] = app
			}

			return app, nil
		}
	case 1:
		{
			if apiContent.OutPutEncoder == "" || apiContent.OutPutEncoder == "origin" {
				step := apiContent.Steps[0]
				balance := step.Balance
				if cfg.Balance != "" {
					balance = cfg.Balance
				}
				balanceK, _ := url.QueryUnescape(balance)
				key := fmt.Sprintf("StaticApp:%d:%s", cfg.ID, balanceK)
				app, has := f.cache[key]
				if !has {
					app = NewDefaultApplication(apiContent, balance)
					f.cache[key] = app
				}

				return app, nil
			}
		}
		fallthrough
	default:
		{
			key := fmt.Sprintf("LayerApp:%d", cfg.ID)
			app, has := f.cache[key]
			if !has {
				app = NewLayerApplication(apiContent)
				f.cache[key] = app
			}
			return app, nil
		}
	}

	return nil, nil
}
