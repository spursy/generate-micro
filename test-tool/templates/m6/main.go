package main

import (
	"os"
	"text/template"
)

func SayHello() string{
	return "hello world"
}

func SayYourName(name string) string {
	return "my name is : " + name
}

func main() {
	funcMap := template.FuncMap{
		"SayHello": SayHello,
		"SayYouName": SayYourName,
	}
	name := "boss"
	tml, err := template.New("test").Funcs(funcMap).Parse("{{SayHello}}\n{{SayYouName .}}\n")
	if err != nil {
		panic(err)
	}
	err = tml.Execute(os.Stdout, name)
	if err != nil {
		panic(err)
	}
}