package updater

type updateHandlerExec struct {
	name     string
	priority int
	*updateHandler
}

type handlerSlice []*updateHandlerExec

func (p handlerSlice) Len() int { // 重写 Len() 方法
	return len(p)
}
func (p handlerSlice) Swap(i, j int) { // 重写 Swap() 方法
	p[i], p[j] = p[j], p[i]
}
func (p handlerSlice) Less(i, j int) bool { // 重写 Less() 方法， 从小到大排序
	return p[i].priority < p[j].priority
}
