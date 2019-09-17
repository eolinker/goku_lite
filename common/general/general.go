package general

type InitFunc func() error

var (
	_initFunc  []InitFunc
	_laterFunc []InitFunc
)

func RegeditInit(fn InitFunc) {

	_initFunc = append(_initFunc, fn)
}
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

func RegeditLater(fn InitFunc) {
	_laterFunc = append(_laterFunc, fn)
}
