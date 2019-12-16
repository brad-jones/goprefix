package prefixer

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/brad-jones/goerr"
)

type Prefixer struct {
	prefix string
}

type PrefixWriter struct {
	p       *Prefixer
	scanner *bufio.Scanner
}

func New(prefix string) *Prefixer {
	return &Prefixer{
		prefix: prefix,
	}
}

func (p *Prefixer) ReadFrom(reader io.Reader) *PrefixWriter {
	return &PrefixWriter{
		p:       p,
		scanner: bufio.NewScanner(reader),
	}
}

func (pw *PrefixWriter) WriteTo(writer io.Writer) error {
	for pw.scanner.Scan() {
		_, err := pw.p.Printf(writer, "%s", pw.scanner.Text())
		if err != nil {
			return goerr.Wrap(err)
		}
	}
	if err := pw.scanner.Err(); err != nil {
		return goerr.Wrap(err)
	}
	return nil
}

func (p *Prefixer) Sprintf(format string, args ...interface{}) string {
	return fmt.Sprintln(p.prefix + strings.TrimSpace(fmt.Sprintf(format, args...)) + "\r")
}

func (p *Prefixer) Printf(writer io.Writer, format string, args ...interface{}) (n int, err error) {
	return fmt.Fprint(writer, p.Sprintf(format, args...))
}
