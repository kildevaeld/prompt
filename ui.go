package prompt

import "io"

type CliUI struct {
	Theme *Theme
	Cursor
	writer io.Writer
}

func (c *CliUI) Password(msg string) string {
	password := &PasswordView{
		Label: msg,
		theme: c.Theme,
	}
	password.Render()

	return password.Value
}

func (c *CliUI) Confirm(msg string) bool {
	confirm := ConfirmView{
		Label: msg,
		theme: c.Theme,
	}

	confirm.Render()

	return confirm.Value
}

func (c *CliUI) List(msg string, choices []string) string {
	list := ListView{
		Label:   msg,
		theme:   c.Theme,
		Choices: choices,
	}

	list.Render()

	return list.Value
}

func (c *CliUI) Form(fields []Field, v ...interface{}) map[string]interface{} {
	form := NewForm(c.Theme, fields)
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
		writer: DefaultTheme.writer,
		Theme:  DefaultTheme,
		Cursor: Cursor{
			writer: DefaultTheme.writer,
		},
	}

}
