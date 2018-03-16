package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/urfave/cli"
)

func Readme(c *cli.Context) error {
	readme := "README.md"

	if Exists(readme) {
		fmt.Println("README.md already exists.")
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

***DEMO:***

## :rocket: Features

- Feature
- ...

## :zap: Requirement

- Requirement
- ...

## :mag: Usage

1. Usage
1. ...

## :package: Installation

$ git clone https://github.com/%s/%s

## :bulb: Anything Else

Anything

## :bust_in_silhouette: Author

[@%s](https://twitter.com/%s)

## :credit_card: License

- [MIT](./LICENSE)`,
		repo, user, repo, user, user)

	_, err = os.Create(readme)
	if err != nil {
		return err
	}
	ioutil.WriteFile(readme, []byte(text), 0666)
	return err
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
