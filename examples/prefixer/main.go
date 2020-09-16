package main

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/brad-jones/goprefix/v2/pkg/prefixer"
)

func main() {
	p1 := prefixer.New("foo | ")

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

	go func() { p1.ReadFrom(stdOutPipe1).WriteTo(os.Stdout) }()
	go func() { p1.ReadFrom(stdErrPipe1).WriteTo(os.Stderr) }()

	if err := cmd1.Wait(); err != nil {
		panic(err)
	}
}

func pingArg() string {
	if runtime.GOOS == "windows" {
		return "-n"
	}
	return "-c"
}
