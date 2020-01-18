package diting

//
//type SummariesProxy struct {
//	ConstLabelsProxy
//	opt *SummaryOpts
//	locker sync.RWMutex
//	summaries Summaries
//}
//
//func newSummariesProxy(opt *SummaryOpts) *SummariesProxy {
//	return &SummariesProxy{
//		ConstLabelsProxy:ConstLabelsProxy(opt.ConstLabels),
//		opt:       opt,
//		locker:    sync.RWMutex{},
//		summaries: nil,
//	}
//}
//
//func (s *SummariesProxy) Refresh(factories Factories) {
//
//	summaries, _ := factories.NewSummary(s.opt)
//	s.locker.Lock()
//	s.summaries = summaries
//	s.locker.Unlock()
//}
//
//func (s *SummariesProxy) Observe(value float64, labels Labels) {
//	s.compile(labels)
//	s.locker.RLock()
//	summaries :=s.summaries
//	s.locker.RUnlock()
//	summaries.Observe(value,labels)
//}
//
