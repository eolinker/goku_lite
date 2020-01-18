package node

import "errors"

var (
	ErrorReadRegisterResultTimeOut = errors.New("read register result timeout")
	ErrorNeedReadRegisterResult    = errors.New("need register-result but not")
	ErrorConsoleRefuse             = errors.New("console refuse")
)
