package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	relativePath := "copy"
	_ = filepath.Walk("./template", func(path string, info fs.FileInfo, err error) error {
		targetPath := strings.Replace(path, "tt", relativePath, -1)
		if info.IsDir() == true {
			return nil
		}

		targetPath = filepath.Join("./", targetPath)

		// 获取文件文件的文件夹路径
		dir := filepath.Dir(targetPath)
		fmt.Printf("dir is %v\n", dir)
		// 创建文件夹
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Println(err)
			}
		}

		// 创建文件
		targetFile, _ := os.Create(targetPath)
		srcFile, _ := os.Open(path)

		ret, err := io.Copy(targetFile, srcFile)
		fmt.Printf("ret --- %v, err --- %v\n", ret, err)
		return nil
	})
}