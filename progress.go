package prompt

import (
	ascii "github.com/kildevaeld/go-ascii"
	tm "github.com/kildevaeld/prompt/terminal"
)

type Progress struct {
	Msg        string
	Theme      *tm.Theme
	ErrorMsg   string
	SuccessMsg string
	msgLen     int
}

func (p *Progress) Done(msg string) {
	p.Theme.Cursor.Show().Backward(p.msgLen)
	p.Theme.Printf("%s%s %s\n", ascii.EraseLine, p.Msg, msg)
}

func (p *Progress) Update(msg string) {
	p.Theme.Cursor.Backward(p.msgLen)
	p.msgLen = p.Theme.Printf("%s%s %s", ascii.EraseLine, p.Msg, p.Theme.HighlightForeground.Color(msg))
}

func (p *Progress) Run(fn func(func(str string)) error) error {
	p.Theme.Cursor.Hide()

	err := fn(p.Update)

	if err != nil {
		p.Done(p.Theme.Error.Color(p.ErrorMsg))
	} else {
		p.Done(p.Theme.Success.Color(p.SuccessMsg))
	}

	return err
}

func NewProgress(msg string, fn func(func(str string)) error) error {

	p := &Progress{
		Msg:        msg,
		Theme:      tm.DefaultTheme,
		ErrorMsg:   "error",
		SuccessMsg: "ok",
	}

	return p.Run(fn)
}
