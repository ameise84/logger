package logger

import (
	"fmt"
	"github.com/ameise84/go_pool"
	"github.com/mattn/go-colorable"
	"io"
	"os"
	"time"
)

var (
	exitHook FatalExitHook
)

func init() {
	exitHook = os.Exit
}

func newWriter() *writer {
	w := &writer{
		cmdW:       colorable.NewColorableStdout(),
		toCmdLevel: LevelTrace,
	}
	w.runner = go_pool.NewGoRunner(w, "log runner",
		go_pool.DefaultOptions().
			SetSimCount(1).
			SetCacheMode(true, 256).
			SetBlock(true, nil),
	)
	return w
}

type writer struct {
	toCmdLevel Level //输出到控制台的级别
	isSync     bool
	cmdW       io.Writer
	file       io.Writer
	logHook    LogHook
	runner     go_pool.GoRunner
}

func (w *writer) OnPanic(err error) {
	w.toConsole(LevelFatal, err.Error())
	w.toFile(err.Error())
}

func (w *writer) OnBlock() {
	_, _ = fmt.Fprintf(os.Stdout, warnFormat, time.Now().Format(layout), "log write is block")
}

func (w *writer) wait() {
	w.runner.Wait()
}

func (w *writer) write(level Level, msg string) {
	if level == LevelFatal {
		_, _ = w.runner.SyncRun(w.writeSync, level, msg)
		exitHook(1)
	} else if w.isSync {
		_, _ = w.runner.SyncRun(w.writeSync, level, msg)
	} else {
		_ = w.runner.AsyncRun(w.writeAsync, level, msg)
	}
}

func (w *writer) writeSync(args ...any) (any, error) {
	level := args[0].(Level)
	msg := args[1].(string)
	w.writeTo(level, msg)
	return nil, nil
}

func (w *writer) writeAsync(args ...any) {
	level := args[0].(Level)
	msg := args[1].(string)
	w.writeTo(level, msg)
}

func (w *writer) writeTo(level Level, msg string) {
	w.toConsole(level, msg)
	if w.logHook != nil || w.file != nil {
		w.toFile(msg)
		w.toHook(level, msg)
	}
}

// 发送到控制台
func (w *writer) toConsole(level Level, msg string) {
	if level >= w.toCmdLevel {
		_, _ = fmt.Fprintf(w.cmdW, ColorMap[level], msg)
	}
}

// 发送到日志钩子
func (w *writer) toHook(level Level, msg string) {
	if w.logHook != nil {
		w.logHook(level, msg)
	}
}

// 写入到文件
func (w *writer) toFile(msg string) {
	if w.file != nil {
		_, _ = w.file.Write([]byte(msg))
	}
}
