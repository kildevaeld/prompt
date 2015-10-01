package widgets

import (
	"time"

	tm "github.com/kildevaeld/prompt/terminal"
)

type View interface {
	Render() error
	SetTheme(theme *tm.Theme)
}

type ConfirmView struct {
	Theme *tm.Theme
	Name  string
	Label string
	Value bool
}

func (c *ConfirmView) Render() {

	if c.Theme == nil {
		c.Theme = tm.DefaultTheme
	}

	label := c.Label
	if label == "" {
		label = c.Name
	}

	c.Theme.Printf("%s [yn]? ", label)
	a, _, _ := tm.GetChar()

	tm.HandleSignals(a)

	ans := string(a)
	if ans == "y" || ans == "Y" {
		c.Value = true
		ans = "yes"
	} else if ans == "n" || ans == "n" {
		c.Value = false
		ans = "no"
	} else {
		c.Theme.Printf("%s%s ", tm.ClearLine, label)
		c.Theme.Highlight("please enter %s(es) or %s(o)", tm.Bold.TextStyle("y"), tm.Bold.TextStyle("n"))

		time.Sleep(1 * time.Second)
		c.Render()
		return
	}
	c.Theme.Highlight("%s\n", ans)
}

func (c *ConfirmView) GetValue() interface{} {
	return c.Value
}

func (c *ConfirmView) GetName() string {
	return c.Name
}

func (c *ConfirmView) SetTheme(theme *tm.Theme) {
	c.Theme = theme
}
