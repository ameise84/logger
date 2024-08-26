package logger

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

var _gLogger *logger

var _gLogLevelStingMap = map[string]Level{
	"trace": LevelTrace,
	"debug": LevelDebug,
	"info":  LevelInfo,
	"warn":  LevelWarn,
	"error": LevelError,
	"fatal": LevelFatal,
}

func init() {
	_gLogger = newLogger()
}

func newLogger() *logger {
	return &logger{
		logLevel: LevelInfo,
		w:        newWriter(),
	}
}

type logger struct {
	logLevel Level //日志等级
	w        *writer
}

func (l *logger) SetSync(isSync bool) {
	l.w.isSync = isSync
}

func (l *logger) SetToCmdLevel(level Level) {
	l.w.toCmdLevel = level
}

func (l *logger) SetLogLevelByName(name string) error {
	lv, ok := _gLogLevelStingMap[name]
	if !ok {
		return errors.New(fmt.Sprintf("not find [%s] log level", name))
	}
	_gLogger.SetLogLevel(lv)
	return nil
}

func (l *logger) SetLogLevel(level Level) {
	l.logLevel = level
}

func (l *logger) SetLogHook(hook LogHook) {
	l.w.logHook = hook
}

func (l *logger) SetFatalExitHook(hook FatalExitHook) {
	exitHook = hook
}

func (l *logger) SetFile(path, file string, maxSize int) error {
	fw, err := newFileWriter(path, file, maxSize)
	if err != nil {
		return err
	}
	l.w.file = fw
	return nil
}

func (l *logger) Wait() {
	l.w.wait()
}

func (l *logger) Trace(arg any) {
	if l.logLevel > LevelTrace {
		return
	}
	l.w.write(LevelTrace, fmt.Sprintf(traceFormat, time.Now().Format(layout), arg))
}

func (l *logger) Debug(arg any) {
	if l.logLevel > LevelDebug {
		return
	}
	l.w.write(LevelDebug, fmt.Sprintf(debugFormat, time.Now().Format(layout), arg))
}

func (l *logger) Info(arg any) {
	if l.logLevel > LevelInfo {
		return
	}
	l.w.write(LevelInfo, fmt.Sprintf(infoFormat, time.Now().Format(layout), arg))
}

func (l *logger) Warn(arg any) {
	if l.logLevel > LevelWarn {
		return
	}
	l.w.write(LevelWarn, fmt.Sprintf(warnFormat, time.Now().Format(layout), arg))
}

func (l *logger) Error(arg any) {
	if l.logLevel > LevelError {
		return
	}
	l.w.write(LevelError, fmt.Sprintf(errorFormat, time.Now().Format(layout), arg))
}

func (l *logger) Fatal(arg any) {
	l.fatal(arg)
}

func (l *logger) TracePrintf(format string, args ...any) {
	if l.logLevel > LevelTrace {
		return
	}
	str := fmt.Sprintf(format, args...)
	l.w.write(LevelTrace, fmt.Sprintf(traceFormat, time.Now().Format(layout), str))
}

func (l *logger) DebugPrintf(format string, args ...any) {
	if l.logLevel > LevelDebug {
		return
	}
	str := fmt.Sprintf(format, args...)
	l.w.write(LevelDebug, fmt.Sprintf(debugFormat, time.Now().Format(layout), str))
}

func (l *logger) InfoPrintf(format string, args ...any) {
	if l.logLevel > LevelInfo {
		return
	}
	str := fmt.Sprintf(format, args...)
	l.w.write(LevelInfo, fmt.Sprintf(infoFormat, time.Now().Format(layout), str))
}

func (l *logger) WarnPrintf(format string, args ...any) {
	if l.logLevel > LevelWarn {
		return
	}
	str := fmt.Sprintf(format, args...)
	l.w.write(LevelWarn, fmt.Sprintf(warnFormat, time.Now().Format(layout), str))
}

func (l *logger) ErrorPrintf(format string, args ...any) {
	if l.logLevel > LevelError {
		return
	}
	str := fmt.Sprintf(format, args...)
	l.w.write(LevelError, fmt.Sprintf(errorFormat, time.Now().Format(layout), str))
}

func (l *logger) FatalPrintf(format string, args ...any) {
	str := fmt.Sprintf(format, args...)
	l.fatal(str)
}

func (l *logger) fatal(arg any) {
	funcName, file, line, _ := runtime.Caller(2)
	s1 := fmt.Sprintf(fatalFormat, time.Now().Format(layout), arg)
	s2 := fmt.Sprintf("%sfunc: %s\t%s:%d\n", s1, runtime.FuncForPC(funcName).Name(), file, line)
	l.w.write(LevelFatal, s2)
}
