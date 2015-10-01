package widgets

import (
	tm "github.com/kildevaeld/prompt/terminal"
)

type PasswordView struct {
	Theme *tm.Theme
	Name  string
	Label string
	Value string
}

func (c *PasswordView) Render() {

	if c.Theme == nil {
		c.Theme = tm.DefaultTheme
	}

	cursor := tm.Cursor{
		Writer: c.Theme,
	}

	label := c.Label
	if label == "" {
		label = c.Name
	}

	c.Theme.Printf("%s ", label)
	x := 0

	buffer := ""

	for {
		a, _, _ := tm.GetChar()
		tm.HandleSignals(a)
		if a == tm.Backspace {
			if x == 0 {
				continue
			}
			c.Theme.Write([]byte("\b \b"))

			x--
			buffer = buffer[0:x]
			continue

		} else if a == tm.Enter {
			c.Value = buffer
			break
		}

		buffer += string(a)
		c.Theme.Write([]byte(c.Theme.Input.Color("*")))

		x++
	}

	cursor.Backward(x)
	str := ""
	for x > 0 {
		str += "*"
		x--
	}

	c.Theme.Highlight("%s\n", str)
}

func (c *PasswordView) GetValue() interface{} {
	return c.Value
}

func (c *PasswordView) GetName() string {
	return c.Name
}

func (c PasswordView) SetTheme(theme *tm.Theme) {
	c.Theme = theme
}
