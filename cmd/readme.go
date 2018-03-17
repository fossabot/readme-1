package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

	"github.com/urfave/cli"
)

type Project struct {
	User string
	Repo string
}

func Readme(c *cli.Context) error {
	readme := "README.md"

	if Exists(readme) {
		err := Edit(readme)
		if err != nil {
			fmt.Println("Cannot open editor.")
			return err
		}
		return nil
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	r := regexp.MustCompile(`.*/src/github.com/.+/.+`)
	if !r.MatchString(dir) {
		fmt.Println("Current directory must be GitHub repository.")
		return err
	}
	el := strings.Split(dir, "/")
	size := len(el)
	project := Project{}
	project.User = el[size-2]
	project.Repo = el[size-1]

	text := `# {{.Repo}}
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

Overview

## :memo: Description

Description

***DEMO***

## :package: Installation

$ git clone https://github.com/{{.User}}/{{.Repo}}

## :rocket: Features

- Feature
- ...

## :zap: Requirement

- Requirement
- ...

## :mag: Usage

1. Usage
1. ...

## :bulb: Example

Example

## :white_check_mark: TODO

- [ ] todo
- [ ] ...

## :bust_in_silhouette: Author

[@{{.User}}](https://twitter.com/{{.User}})

## :credit_card: License

- [MIT](./LICENSE) &copy; {{.User}}`

	file, err := os.Create(readme)
	if err != nil {
		return err
	}
	t := template.New("t")
	template.Must(t.Parse(text))
	t.Execute(file, project)

	err = Edit(readme)
	return err
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func Edit(filename string) error {
	editor := os.Getenv("EDITOR")
	cmd := exec.Command(editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("Cannot open editor.")
		return err
	}
	return err
}
