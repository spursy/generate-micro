package main

import (
	"os"
	"text/template"
)

type MyMethod struct {
	Say string
	Name string
}

func (my *MyMethod) SayHello() string {
	return "world"
}

func (my *MyMethod)SayYourName(name string ) string {
	return "my name is : " + name
}

func main() {
	mine := &MyMethod{
		Say: "hello",
		Name: "student",
	}
	tmpl, err := template.New("test").Parse("{{$str1 := .Say}}{{$str2 := .SayHello}}{{$str3 := .SayYourName .Name}}{{$str1}} {{$str2}}\n{{$str3}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, mine)
	if err != nil {
		panic(err)
	}
}