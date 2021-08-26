package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count uint
}

func main() {
	sweaters := Inventory{"Wool", 17}
	tml, err := template.New("test").Parse("{{.Count}} of {{.Material}}\n")
	if err != nil {
		panic(err)
	}
	err = tml.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}