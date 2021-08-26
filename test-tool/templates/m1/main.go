package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	fn := template.FuncMap{
		"title": func(s string) string {
			return strings.ReplaceAll(strings.Title(s), "-", "")
		},
		"dehyphen": func(s string) string {
			return strings.ReplaceAll(s, "-", "")
		},
		"lower": func(s string) string {
			return strings.ToLower(s)
		},
	}
	f, _ := os.Create("./data")
	t, _ := template.New("f").Funcs(fn).Parse("")
	_ = t.Execute(f, "")
}