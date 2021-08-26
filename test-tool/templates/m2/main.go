package main

import (
	"os"
	"text/template"
)

func main() {
	str := "world"
	tml, err := template.New("test").Parse("hello, {{.}}\n")
	if err != nil {
		panic(err)
	}
	err = tml.Execute(os.Stdout, str)
	if err != nil {
		panic(err)
	}
}