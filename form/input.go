package form

import (
	"github.com/kildevaeld/go-widgets"
	tm "github.com/kildevaeld/prompt/terminal"
)

type Input struct {
	Theme *tm.Theme
	Name  string
	widgets.Input
	Message     string
	Validations Validations
}

func (c *Input) Run() {

	if c.Theme == nil {
		c.Theme = tm.DefaultTheme
	}

	c.Config = themeToWidgetConfig(c.Theme)

	label := c.Message
	if label == "" {
		label = c.Name
	}
	c.Input.Message = label

	c.Input.Run()
}

func (c *Input) GetValue() interface{} {
	return c.Value
}

func (c *Input) GetName() string {
	return c.Name
}

func (c *Input) SetTheme(theme *tm.Theme) {
	c.Theme = theme
}
