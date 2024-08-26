package example

import (
	"errors"
	"github.com/ameise84/logger"
	"strconv"
	"sync"
	"testing"
	"time"
)

func call(n int) error {
	return errors.New("something error in call3 -> param:" + strconv.Itoa(n))
}

func TestLogger(t *testing.T) {
	logger.SetLogLevel(logger.LevelTrace)
	_ = logger.SetFile(".", "app", 1)
	logger.SetLogHook(hook)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < 2; i++ {
			logger.TracePrintf("this is trace[%v]", i)
			logger.Debug("this is debug")
			logger.Info("succeeded")
			logger.Warn("warn!!!")
			logger.Error("error")
			logger.Fatal(call(2))
		}
		wg.Done()
	}()
	wg.Wait()
	logger.Wait()
}

func hook(level logger.Level, s string) {
	mp := map[string]any{}
	if level == logger.LevelError {
		mp["error"] = s
	}
	panic("error")
}
