package general

//InitFunc init func
type InitFunc func() error

var (
	_initFunc  []InitFunc
	_laterFunc []InitFunc
)

//RegeditInit 初始化注册
func RegeditInit(fn InitFunc) {

	_initFunc = append(_initFunc, fn)
}

//General general
func General() error {
	for _, fn := range _initFunc {
		if err := fn(); err != nil {
			return err
		}
	}
	for _, fn := range _laterFunc {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

//RegeditLater regedit later
func RegeditLater(fn InitFunc) {
	_laterFunc = append(_laterFunc, fn)
}
