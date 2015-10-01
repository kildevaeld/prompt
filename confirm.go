package prompt

import "time"

type View interface {
	Render() error
	SetTheme(theme *Theme)
}

type ConfirmView struct {
	theme *Theme
	Name  string
	Label string
	Value bool
}

func (c *ConfirmView) Render() {

	if c.theme == nil {
		c.theme = DefaultTheme
	}

	label := c.Label
	if label == "" {
		label = c.Name
	}

	c.theme.Printf("%s [yn]? ", label)
	a, _, _ := getChar()

	handleSignals(a)

	ans := string(a)
	if ans == "y" || ans == "Y" {
		c.Value = true
		ans = "yes"
	} else if ans == "n" || ans == "n" {
		c.Value = false
		ans = "no"
	} else {
		c.theme.Printf("%s%s ", ClearLine, label)
		c.theme.Highlight("please enter %s(es) or %s(o)", Bold.TextStyle("y"), Bold.TextStyle("n"))

		time.Sleep(1 * time.Second)
		c.Render()
		return
	}
	c.theme.Highlight("%s\n", ans)
}

func (c *ConfirmView) GetValue() interface{} {
	return c.Value
}

func (c *ConfirmView) GetName() string {
	return c.Name
}

func (c *ConfirmView) SetTheme(theme *Theme) {
	c.theme = theme
}
