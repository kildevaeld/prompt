package prompt

import (
	"time"

	ascii "github.com/kildevaeld/go-ascii"
	tm "github.com/kildevaeld/prompt/terminal"
	"github.com/tj/go-spin"
)

type Process struct {
	Msg        string
	Theme      *tm.Theme
	done       chan bool
	msgLen     int
	ErrorMsg   string
	SuccessMsg string
}

func (p *Process) Start() {
	p.Theme.Cursor.Hide()

	p.done = make(chan bool)

	ticker := time.NewTicker(100 * time.Millisecond)
	s := spin.New()

	go func() {
	loop:
		for {

			select {
			case <-p.done:
				ticker.Stop()
				break loop
			case <-ticker.C:
				p.update(s.Next())
			}

		}

		close(p.done)

	}()

}

func (p *Process) Run(fn func() error) error {
	p.Start()
	err := fn()

	if err != nil {
		p.Done(p.Theme.Error.Color(p.ErrorMsg))
	} else {
		p.Done(p.Theme.Success.Color(p.SuccessMsg))
	}
	return err
}

func (p *Process) update(msg string) {
	p.Theme.Cursor.Backward(p.msgLen)
	p.msgLen = p.Theme.Printf("%s%s %s", ascii.EraseLine, p.Msg, p.Theme.HighlightForeground.Color(msg))
}

func (p *Process) Done(msg string) {
	p.done <- true
	p.Theme.Cursor.Show().Backward(p.msgLen)
	p.Theme.Printf("%s%s %s\n", ascii.EraseLine, p.Msg, msg)
}

func NewProcess(msg string, fn func() error) error {

	p := &Process{
		Msg:        msg,
		Theme:      tm.DefaultTheme,
		ErrorMsg:   "error",
		SuccessMsg: "ok",
	}

	return p.Run(fn)
}
