package main

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/brad-jones/goprefix/v2/pkg/colorchooser"
	"github.com/brad-jones/goprefix/v2/pkg/prefixer"
)

func main() {
	p1 := prefixer.New(colorchooser.Sprint("ip1") + " | ")
	cmd1 := exec.Command("ping", pingArg(), "4", "127.0.0.1")
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

	p2 := prefixer.New(colorchooser.Sprint("ip2") + " | ")
	cmd2 := exec.Command("ping", pingArg(), "4", "127.0.0.2")
	stdOutPipe2, err := cmd2.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdOutPipe2.Close()
	stdErrPipe2, err := cmd2.StderrPipe()
	if err != nil {
		panic(err)
	}
	defer stdErrPipe2.Close()
	if err := cmd2.Start(); err != nil {
		panic(err)
	}

	go func() { p1.ReadFrom(stdOutPipe1).WriteTo(os.Stdout) }()
	go func() { p1.ReadFrom(stdErrPipe1).WriteTo(os.Stderr) }()
	go func() { p2.ReadFrom(stdOutPipe2).WriteTo(os.Stdout) }()
	go func() { p2.ReadFrom(stdErrPipe2).WriteTo(os.Stderr) }()

	if err := cmd1.Wait(); err != nil {
		panic(err)
	}
	if err := cmd2.Wait(); err != nil {
		panic(err)
	}
}

func pingArg() string {
	if runtime.GOOS == "windows" {
		return "-n"
	}
	return "-c"
}
