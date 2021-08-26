package main

import (
	"os"
	"text/template"
)

type MyMethod struct {
	Name string
}

func (my *MyMethod)SayHello() string {
	return "hello world"
}

func (my *MyMethod)SayYourName(name string) string {
	return "my name is : " + name
}

func main() {
	mine := &MyMethod{Name: "Boss"}
	tmpl, err := template.New("test").Parse("{{.SayHello}}\n{{.SayYourName .Name}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, mine)
	if err != nil {
		panic(err)
	}
}