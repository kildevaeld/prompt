package prompt

import (
	"fmt"
	"os"
	"time"

	tm "github.com/kildevaeld/prompt/terminal"
)

type Progress struct {
	Msg string
}

func (p *Progress) Done(msg string) {
	p.Update(msg + "\n")
}

func (p *Progress) Update(msg string) {
	fmt.Printf("\r\033[0K\033[90m%s %s", p.Msg, tm.Cyan.Color(msg))
}

func NewProgress(msg string, fn func(func(str string)) error) error {

	p := &Progress{
		Msg: msg,
	}
	os.Stdout.Write([]byte(tm.HideCursor))
	err := fn(p.Update)

	if err != nil {
		p.Done(tm.Red.Color("error"))
	} else {
		p.Done(tm.Green.Color("ok"))
	}
	time.Sleep(300 * time.Millisecond)
	os.Stdout.Write([]byte(tm.ShowCursor))
	return err
}
