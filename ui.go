package prompt

import (
	"io"

	"github.com/kildevaeld/prompt/terminal"
	"github.com/kildevaeld/prompt/widgets"
)

type CliUI struct {
	Theme *terminal.Theme
	terminal.Cursor
	writer io.Writer
}

func (c *CliUI) Password(msg string) string {
	password := &widgets.PasswordView{
		Label: msg,
		Theme: c.Theme,
	}
	password.Render()

	return password.Value
}

func (c *CliUI) Confirm(msg string) bool {
	confirm := &widgets.ConfirmView{
		Label: msg,
		Theme: c.Theme,
	}

	confirm.Render()

	return confirm.Value
}

func (c *CliUI) List(msg string, choices []string) string {
	list := &widgets.ListView{
		Label:   msg,
		Theme:   c.Theme,
		Choices: choices,
	}

	list.Render()

	return list.Value
}

func (c *CliUI) Form(fields []widgets.Field, v ...interface{}) map[string]interface{} {
	form := widgets.NewForm(c.Theme, fields)
	form.Render()

	if len(v) > 0 {
		form.GetValue(v[0])
	}

	return form.Value
}

func (c *CliUI) Clear() {
	c.writer.Write([]byte("\033[2J"))
}

func (c *CliUI) Save() {

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
