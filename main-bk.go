package main

import (
	"bufio"
	"fmt"
	"git.5th.im/long-bridge-algo/golang/micro-gen-go/helper"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Printf("出现异常:%v，\n请联系 yang.yang@longbridge.sg\n", err)
		}
	}()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入项目名：")
	projectName, err := reader.ReadString('\n')

	_ = filepath.Walk("./template", func(path string, info fs.FileInfo, err error) error {
		targetPath := strings.Replace(path, "template", projectName, -1)
		if info.IsDir() == true {
			return nil
		}
		// 执行的目录
		exPath, err := helper.GetExPath()
		if err != nil {
			fmt.Printf("出现异常:%v，\n请联系 yang.yang@longbridge.sg\n", err)
			return nil
		}
		targetPath = filepath.Join(exPath, targetPath)

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