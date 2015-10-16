package form

import (
	"github.com/kildevaeld/go-widgets"
	tm "github.com/kildevaeld/prompt/terminal"
)

type Password struct {
	Theme *tm.Theme
	Name  string
	widgets.Password
	Message string
}

func (c *Password) Run() {

	if c.Theme == nil {
		c.Theme = tm.DefaultTheme
	}

	c.Config = themeToWidgetConfig(c.Theme)

	label := c.Message
	if label == "" {
		label = c.Name
	}
	c.Password.Message = label

	c.Password.Run()
}

func (c *Password) GetValue() interface{} {
	return c.Value
}

func (c *Password) GetName() string {
	return c.Name
}

func (c *Password) SetTheme(theme *tm.Theme) {
	c.Theme = theme
}
