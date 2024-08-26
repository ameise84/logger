package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func newFileWriter(path, fileName string, maxSize int) (*fileWriter, error) {
	baseName := filepath.Join(path, fileName)
	err := os.MkdirAll(path, 0666)
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(baseName+"_last"+logExt, os.O_CREATE|os.O_WRONLY|os.O_SYNC|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	w := &fileWriter{
		baseName: baseName,
		file:     file,
		maxSize:  maxSize * MB,
		size:     int(stat.Size()),
		index:    searchMaxIndexFiles(path, fileName),
	}
	return w, nil
}

type fileWriter struct {
	baseName string
	file     *os.File
	maxSize  int
	size     int
	index    int
}

func (w *fileWriter) Write(p []byte) (n int, err error) {
	n, err = w.file.Write(p)
	if err != nil {
		fmt.Printf(errorFormat, time.Now().Format(layout), err)
	}
	w.size += n
	if w.size > w.maxSize {
		if err != nil {
			fmt.Printf(errorFormat, time.Now().Format(layout), err)
			return
		}
		err = w.file.Close()
		if err != nil {
			fmt.Printf(errorFormat, time.Now().Format(layout), err)
			return
		}
		oldName := w.file.Name()
		err = os.Rename(oldName, w.baseName+"_"+strconv.Itoa(w.index)+logExt)
		if err != nil {
			fmt.Printf(errorFormat, time.Now().Format(layout), err)
			return
		}
		w.file, err = os.OpenFile(oldName, os.O_CREATE|os.O_WRONLY|os.O_SYNC|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Printf(errorFormat, time.Now().Format(layout), err)
			return
		}
		w.size = 0
		w.index++
	}
	return n, nil
}
