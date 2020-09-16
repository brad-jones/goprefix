# goprefix

[![PkgGoDev](https://pkg.go.dev/badge/github.com/brad-jones/goprefix/v2)](https://pkg.go.dev/github.com/brad-jones/goprefix/v2)
[![GoReport](https://goreportcard.com/badge/github.com/brad-jones/goprefix/v2)](https://goreportcard.com/report/github.com/brad-jones/goprefix/v2)
[![GoLang](https://img.shields.io/badge/golang-%3E%3D%201.15.1-lightblue.svg)](https://golang.org)
![.github/workflows/main.yml](https://github.com/brad-jones/goprefix/workflows/.github/workflows/main.yml/badge.svg?branch=v2)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)
[![KeepAChangelog](https://img.shields.io/badge/Keep%20A%20Changelog-1.0.0-%23E05735)](https://keepachangelog.com/)
[![License](https://img.shields.io/github/license/brad-jones/goprefix.svg)](https://github.com/brad-jones/goprefix/blob/v2/LICENSE)

A simple library to help prefix streams of text, useful to create "docker-compose" like interfaces.

The `prefixer` will read a stream and prefix each line with the given prefix.

Given a string the `colorchooser` will return a string colored with a random
16-bit color, given the same string again, it will be returned in the same color.

_Looking for v1, see the [master branch](https://github.com/brad-jones/goprefix/tree/master)_

## Quick Start

`go get -u github.com/brad-jones/goprefix/v2/pkg/...`

```go
package main

import (
	"os"
	"os/exec"

	"github.com/brad-jones/goprefix/v2/pkg/colorchooser"
	"github.com/brad-jones/goprefix/v2/pkg/prefixer"
)

func main() {
	p1 := prefixer.New(colorchooser.Sprint("ping 1.1.1.1") + " | ")
	cmd1 := exec.Command("ping", "-c", "4", "1.1.1.1")

	stdOutPipe1, err := cmd1.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdOutPipe1.Close()
	
	stdErrPipe1, err := cmd1.StderrPipe()
	if err != nil {
		panic(err)
	}
	defer stdErrPipe1.Close()
	
	if err := cmd1.Start(); err != nil {
		panic(err)
	}

	go func() { p1.ReadFrom(stdOutPipe1).WriteTo(os.Stdout) }()
	go func() { p1.ReadFrom(stdErrPipe1).WriteTo(os.Stderr) }()

	if err := cmd1.Wait(); err != nil {
		panic(err)
	}
}

```

_Also see further working examples under: <https://github.com/brad-jones/goexec/tree/v2/examples>_
