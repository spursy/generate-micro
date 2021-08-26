package main

import (
	"fmt"
	"strings"
)

func main() {
	absPath := "tt"
	for _, item := range []int64{1, 2} {
		fmt.Println(item)
		ret := strings.Replace("./template/t1/asa", "./template", absPath, -1)
		fmt.Printf(" dest path: %v\n", ret)
	}
}