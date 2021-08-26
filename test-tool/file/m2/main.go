package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	f := filepath.Join("./test", "test1.txt")
	dir := filepath.Dir(f)
	fmt.Println(dir)

	// 创建文件夹
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println(err)
		}
	}

	// 创建文件
	targetFile, _ := os.Create(f)
	srcFile, _ := os.Open("./test.txt")

	ret, err := io.Copy(targetFile, srcFile)
	fmt.Printf("ret --- %v, err --- %v\n", ret, err)
}