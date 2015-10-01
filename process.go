package prompt

import (
	"fmt"
	"os"
	"time"

	tm "github.com/kildevaeld/prompt/terminal"
	"github.com/tj/go-spin"
)

type Process struct {
	Msg  string
	done chan bool
}

func (p *Process) Start() {
	os.Stdout.Write([]byte(tm.HideCursor))
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

func (p *Process) update(msg string) {
	fmt.Printf("\r%s%s %s\r", tm.Gray, p.Msg, tm.Cyan.Color(msg))

}

func (p *Process) Done(msg string) {
	p.done <- true
	os.Stdout.Write([]byte(tm.ShowCursor))
	fmt.Printf("\r%s%s %s\n", tm.Gray, p.Msg, msg)

}

func NewProcess(msg string, fn func() error) error {

	p := &Process{Msg: msg}

	p.Start()

	err := fn()
	time.Sleep(300 * time.Millisecond)
	if err != nil {
		p.Done(tm.Red.Color("error"))
	} else {
		p.Done(tm.Green.Color("ok"))
	}
	return err
}
