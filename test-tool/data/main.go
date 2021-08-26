package main

import (
	"fmt"
	"git.5th.im/long-bridge-algo/golang/micro-gen-go/test-tool/data/bindata"
)

func main() {
	fmt.Println("0000000")
	data, err := bindata.Asset("data")
	if err != nil {
		fmt.Println("111111")
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", string(data))
}