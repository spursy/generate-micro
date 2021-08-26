package main

import (
	"fmt"
	"go/build"
)

func main() {
	goPath := build.Default.GOPATH
	fmt.Println(goPath)
}