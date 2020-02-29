package goku_plugin

import "fmt"

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})

	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
}


// Debug logs a message at level Debug on the standard _logger.
func Debug(args ...interface{}) {
	if _logger == nil{
		printLevel("[DEBUG]",args...)
		return
	}
	_logger.Debug(args...)
}


// Info logs a message at level Info on the standard _logger.
func Info(args ...interface{}) {
	if _logger == nil{
		printLevel("[INFO]",args...)
		return
	}
	_logger.Info(args...)
}

// Warn logs a message at level Warn on the standard _logger.
func Warn(args ...interface{}) {
	if _logger == nil{
		printLevel("[WARN]",args...)
		return
	}
	_logger.Warn(args...)
}

// Warning logs a message at level Warn on the standard _logger.
func Warning(args ...interface{}) {
	if _logger == nil{
		printLevel("[WARNING]",args...)
		return
	}
	_logger.Warning(args...)
}

// Error logs a message at level Error on the standard _logger.
func Error(args ...interface{}) {
	if _logger == nil{
		printLevel("[ERROR]",args...)
		return
	}
	_logger.Error(args...)
}

// Debugf logs a message at level Debug on the standard _logger.
func Debugf(format string, args ...interface{}) {
	if _logger == nil{
		printLevelF(format,"[DEBUG] ",args...)
		return
	}
	_logger.Debugf(format, args...)
}

// Infof logs a message at level Info on the standard _logger.
func Infof(format string, args ...interface{}) {
	if _logger == nil{
		printLevelF(format,"[INFO] ",args...)
		return
	}
	_logger.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard _logger.
func Warnf(format string, args ...interface{}) {
	if _logger == nil{
		printLevelF(format,"[WARN] ",args...)
		return
	}
	_logger.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard _logger.
func Warningf(format string, args ...interface{}) {
	if _logger == nil{
		printLevelF(format,"[WARNING] ",args...)
		return
	}
	_logger.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard _logger.
func Errorf(format string, args ...interface{}) {
	if _logger == nil{
		printLevelF(format,"[ERROR] ",args...)
		return
	}
	_logger.Errorf(format, args...)
}

func printLevel(level string,args ...interface{}){
	vs:=make([]interface{},0,len(args)+1)
	vs = append(vs,level)
	vs = append(vs,args...)
	fmt.Println(vs...)
}
func printLevelF(format,level string,args ...interface{}){
	if format[len(format)-1] !='\n'{
		format = format+"\n"
	}
	fmt.Printf(level+format,args...)
}