package main

import (
	"bufio"
	"fmt"
	tmpl "git.5th.im/long-bridge-algo/golang/micro-gen-go/template"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type config struct {
	// foo
	Alias string
	// github.com/micro/foo
	Dir string
	// $GOPATH/src/github.com/micro/foo
	GoDir string
	// $GOPATH
	GoPath string
	// UseGoPath
	UseGoPath bool
	// Files
	Files []file
	// Comments
	Comments []string
}

type file struct {
	Path string
	Tmpl string
}

func Create(c config) error{
	// check if dir exists
	if _, err := os.Stat(c.Dir); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", c.Dir)
	}

	// write the files
	for _, file := range c.Files {
		f := filepath.Join(c.Dir, file.Path)
		dir := filepath.Dir(f)

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}
		if err := write(c, f, file.Tmpl); err != nil {
			return err
		}
	}
	return nil
}

func write(c config, file, tmpl string) error {
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

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.New("f").Funcs(fn).Parse(tmpl)
	if err != nil {
		return err
	}

	return t.Execute(f, c)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入项目名：")
	projectName, err := reader.ReadString('\n')
	if err != nil {
		// ToDo
		return
	}

	c := config{
		Alias:     projectName,
		Comments:  nil,
		Dir:       projectName,
		GoDir:     nil,
		GoPath:    nil,
		UseGoPath: false,
		Files: []file{
			{"README.md", tmpl.ReadmeSRV},
			{".gitlab-ci.yml", tmpl.GitLabCiSRV},
		},
	}
	err = Create(c)
	if err != nil {
		fmt.Printf("The error is %v\n", err)
	}
}