package dao

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ReadTextFile(p string, absolute bool) (res string, err error) {
	absPath := filepath.Join(BasePath, p)
	if absolute {
		absPath = p
	}

	// 检查文件存在
	_, err = os.Stat(absPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", errors.New(fmt.Sprintf("file %s does not exist", absPath))
		}
		return "", err
	}

	// 打开文件
	file, err := os.Open(absPath)
	defer file.Close()
	if err != nil {
		return "", err
	}

	// 读取数据
	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), err
}
