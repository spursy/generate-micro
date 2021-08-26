package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aliba-ba"
	ret := strings.ReplaceAll(strings.Title(s), "-", "")
	fmt.Println(ret)
}