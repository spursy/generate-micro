package helper

import (
	"os"
	"path/filepath"
)

/*
获取执行的目录
*/
func GetExPath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}