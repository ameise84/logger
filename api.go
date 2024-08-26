package logger

import (
	"fmt"
)

type Level = int

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

type LogHook func(Level, string)

type FatalExitHook func(int)

type Log interface {
	Trace(any)
	Debug(any)
	Info(any)
	Warn(any)
	Error(any)
	Fatal(any)
	TracePrintf(string, ...any)
	DebugPrintf(string, ...any)
	InfoPrintf(string, ...any)
	WarnPrintf(string, ...any)
	ErrorPrintf(string, ...any)
	FatalPrintf(string, ...any)
}

type Logger interface {
	Log
	SetSync(bool)
	SetLogLevelByName(string) error
	SetLogLevel(Level)
	SetLogHook(LogHook)
	SetFatalExitHook(FatalExitHook)
	SetFile(path, file string, maxSize int) error // the file in: path/file_x.log  ,maxSize is MB
	SetToCmdLevel(Level)                          //是否输出到控制台
	Wait()
}

func NewLogger() Logger {
	return newLogger()
}

func DefaultLogger() Logger {
	return _gLogger
}

func SetSync(isSync bool) {
	_gLogger.SetSync(isSync)
}

func SetToCmdLevel(level Level) {
	_gLogger.SetToCmdLevel(level)
}

func SetLogLevelByName(name string) error {
	return _gLogger.SetLogLevelByName(name)
}

func SetLogLevel(level Level) {
	_gLogger.SetLogLevel(level)
}

func SetLogHook(hook LogHook) {
	_gLogger.SetLogHook(hook)
}

func SetFatalExitHook(hook FatalExitHook) {
	_gLogger.SetFatalExitHook(hook)
}

func SetFile(path, file string, maxSize int) error {
	return _gLogger.SetFile(path, file, maxSize)
}

func Wait() {
	_gLogger.Wait()
}

func Trace(msg any) {
	_gLogger.Trace(msg)
}

func Debug(msg any) {
	_gLogger.Debug(msg)
}

func Info(msg any) {
	_gLogger.Info(msg)
}

func Warn(msg any) {
	_gLogger.Warn(msg)
}

func Error(msg any) {
	_gLogger.Error(msg)
}

func Fatal(msg any) {
	_gLogger.fatal(msg)
}

func TracePrintf(format string, msg ...any) {
	_gLogger.TracePrintf(format, msg...)
}

func DebugPrintf(format string, msg ...any) {
	_gLogger.DebugPrintf(format, msg...)
}

func InfoPrintf(format string, msg ...any) {
	_gLogger.InfoPrintf(format, msg...)
}

func WarnPrintf(format string, msg ...any) {
	_gLogger.WarnPrintf(format, msg...)
}

func ErrorPrintf(format string, msg ...any) {
	_gLogger.ErrorPrintf(format, msg...)
}

func FatalPrintf(format string, msg ...any) {
	str := fmt.Sprintf(format, msg...)
	_gLogger.fatal(str)
}
