package form

import (
	"github.com/kildevaeld/go-widgets"
	tm "github.com/kildevaeld/prompt/terminal"
)

type List struct {
	Theme *tm.Theme
	Name  string
	widgets.List
	Message string
	Choices []string
}

func (c *List) Run() {

	if c.Theme == nil {
		c.Theme = tm.DefaultTheme
	}

	c.Config = themeToWidgetConfig(c.Theme)

	label := c.Message
	if label == "" {
		label = c.Name
	}
	c.List.Message = label
	c.List.Choices = c.Choices
	c.List.Run()
}

func (c *List) GetValue() interface{} {
	return c.Value
}

func (c *List) GetName() string {
	return c.Name
}

func (c *List) SetTheme(theme *tm.Theme) {
	c.Theme = theme
}
