package diting

//Proxy proxy
type Proxy interface {
	Refresh(factories Factories)
}

//ConstLabelsProxy 为了能定义 constLabel 内字段的顺序，将constLabel 改为普通Label
type ConstLabelsProxy Labels

func (p ConstLabelsProxy) compile(labels Labels) {
	if p == nil || labels == nil {
		return
	}
	for k, v := range p {
		if _, has := labels[k]; !has {
			labels[k] = v
		}
	}
}
