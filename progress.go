package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/ttacon/chalk"
)

type Progress struct {
	Msg string
}

func (p *Progress) Done(msg string) {
	p.Update(msg + "\n")
}

func (p *Progress) Update(msg string) {
	fmt.Printf("\r\033[0K\033[90m%s %s", p.Msg, chalk.Cyan.Color(msg))
}

func NewProgress(msg string, fn func(func(str string)) error) error {

	p := &Progress{
		Msg: msg,
	}
	os.Stdout.Write([]byte(HideCursor))
	err := fn(p.Update)

	if err != nil {
		p.Done(chalk.Red.Color("error"))
	} else {
		p.Done(chalk.Green.Color("ok"))
	}
	time.Sleep(300 * time.Millisecond)
	os.Stdout.Write([]byte(ShowCursor))
	return err
}
