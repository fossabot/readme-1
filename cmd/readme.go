package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/urfave/cli"
)

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
	user := el[size-2]
	repo := el[size-1]

	text := fmt.Sprintf(`# %s
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

Overview

## :memo: Description

Description

***DEMO***

## :package: Installation

$ git clone https://github.com/%s/%s

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

[@%s](https://twitter.com/%s)

## :credit_card: License

- [MIT](./LICENSE) &copy; micnncim`,
		repo, user, repo, user, user)

	_, err = os.Create(readme)
	if err != nil {
		return err
	}
	ioutil.WriteFile(readme, []byte(text), 0666)
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
