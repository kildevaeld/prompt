package prompt

import (
	"fmt"
	"io"

	"github.com/kildevaeld/prompt/form"
	"github.com/kildevaeld/prompt/terminal"
)

type CliUI struct {
	Theme *terminal.Theme
	terminal.Cursor
	writer io.Writer
}

func (c *CliUI) Password(msg string) string {
	password := &form.Password{
		Message: msg,
		Theme:   c.Theme,
	}
	password.Run()

	return password.Value
}

func (c *CliUI) Confirm(msg string) bool {
	confirm := &form.Confirm{
		Message: msg,
		Theme:   c.Theme,
	}

	confirm.Run()

	return confirm.Value
}

func (c *CliUI) List(msg string, choices []string) string {
	list := &form.List{
		Message: msg,
		Theme:   c.Theme,
		Choices: choices,
	}

	list.Run()

	return list.Value
}

func (c *CliUI) PaginatedList(msg string, paginate func(page int) []string) string {
	list := &form.PaginatedList{
		Message:  msg,
		Theme:    c.Theme,
		Paginate: paginate,
	}

	list.Run()

	return list.Value
}

func (c *CliUI) Input(msg string) string {
	input := &form.Input{
		Message: msg,
		Theme:   c.Theme,
	}

	input.Run()

	return input.Value

}

func (c *CliUI) Process(msg string, args ...interface{}) *Process {

	process := &Process{
		Theme:      c.Theme,
		Msg:        fmt.Sprintf(msg, args...),
		ErrorMsg:   "error",
		SuccessMsg: "success",
	}

	return process
}

func (c *CliUI) Progress(msg string) *Progress {

	progress := &Progress{
		Theme:      c.Theme,
		Msg:        msg,
		ErrorMsg:   "error",
		SuccessMsg: "success",
	}

	return progress
}

func (c *CliUI) FormWithFields(fields []form.Field, v ...interface{}) map[string]interface{} {
	form := form.NewForm(c.Theme, fields)
	form.Run()

	if len(v) > 0 {
		form.GetValue(v[0])
	}

	return form.Value
}

func (c *CliUI) Form(v interface{}) error {
	return form.FormFromStruct(c.Theme, v)
}

func (c *CliUI) Clear() {
	c.writer.Write([]byte("\033[2J"))
	c.Move(0, 0)
}

func (c *CliUI) Save() {
	terminal.Save()
}

func (c *CliUI) Restore() {
	terminal.Restore()
}

func (c *CliUI) Printf(msg string, args ...interface{}) {
	c.Theme.Printf(msg, args...)
}

func NewUI() *CliUI {

	return &CliUI{
		writer: terminal.DefaultTheme.Writer,
		Theme:  terminal.DefaultTheme,
		Cursor: terminal.Cursor{
			Writer: terminal.DefaultTheme.Writer,
		},
	}

}
