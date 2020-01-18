package internal

//RouterRule RouterRule
type RouterRule interface {
	GetTargets() []int
	Match(arg ...string) bool
}
