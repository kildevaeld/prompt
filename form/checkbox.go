package form

import (
	"github.com/kildevaeld/go-widgets"
	tm "github.com/kildevaeld/prompt/terminal"
)

type Checkbox struct {
	Theme   *tm.Theme
	Name    string
	Message string
	Choices []string
	widgets.Checkbox
}

func (c *Checkbox) Run() {

	if c.Theme == nil {
		c.Theme = tm.DefaultTheme
	}

	c.Config = themeToWidgetConfig(c.Theme)

	label := c.Message
	if label == "" {
		label = c.Name
	}
	c.Checkbox.Message = label
	c.Checkbox.Choices = c.Choices
	c.Checkbox.Run()
}

func (c *Checkbox) GetValue() interface{} {
	return c.Value
}

func (c *Checkbox) GetName() string {
	return c.Name
}

func (c *Checkbox) SetTheme(theme *tm.Theme) {
	c.Theme = theme
}
