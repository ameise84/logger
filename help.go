package logger

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func searchMaxIndexFiles(root string, searchTerm string) int {
	var re = regexp.MustCompile("^" + searchTerm + "_([\\d]+)" + logExt + "$")
	// 遍历目录树
	maxIndex := 0
	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() { // 检查是否是文件
			ss := re.FindStringSubmatch(path)
			if ss != nil {
				index, _ := strconv.Atoi(ss[1])
				maxIndex = max(index+1, maxIndex)
			}
		}
		return nil
	})

	return maxIndex
}
