package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入项目名：")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}