package general

//InitFunc 初始化函数
type InitFunc func() error

var (
	_initFunc  []InitFunc
	_laterFunc []InitFunc
)

//RegeditInit 注册初始化
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

//RegeditLater regeditlater
func RegeditLater(fn InitFunc) {
	_laterFunc = append(_laterFunc, fn)
}
