package form

import (
	"github.com/kildevaeld/go-widgets"
	tm "github.com/kildevaeld/prompt/terminal"
)

func themeToWidgetConfig(theme *tm.Theme) widgets.WidgetConfig {
	return widgets.WidgetConfig{
		MessageColor:   theme.Foreground.Style,
		HighlightColor: theme.HighlightForeground.Style,
		StdinColor:     theme.Input.Style,
		Writer:         theme.Writer,
	}
}

type Confirm struct {
	Theme *tm.Theme
	Name  string
	//Message string
	//Value bool
	Message string
	widgets.Confirm
}

func (c *Confirm) Run() {

	if c.Theme == nil {
		c.Theme = tm.DefaultTheme
	}

	c.Config = themeToWidgetConfig(c.Theme)

	label := c.Message
	if label == "" {
		label = c.Name
	}
	c.Confirm.Message = label

	c.Confirm.Run()
}

func (c *Confirm) GetValue() interface{} {
	return c.Value
}

func (c *Confirm) GetName() string {
	return c.Name
}

func (c *Confirm) SetTheme(theme *tm.Theme) {
	c.Theme = theme
}
