package widgets

import (
	tm "github.com/kildevaeld/prompt/terminal"
)

type ListView struct {
	Name    string
	Label   string
	Value   string
	Choices []string
	Theme   *tm.Theme
}

func (c *ListView) Render() {
	choices := c.Choices
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
	cursor.Hide()
	c.Theme.Printf("%s\n", label)

	for i, s := range choices {
		if i == len(choices)-1 {
			c.highlight_line(s)
		} else {
			c.print_line(s)
		}
		c.Theme.WriteString("\n")

	}
	l := len(choices)

	cursor.Up(1)
	curPos := l - 1
	for {
		a, k, e := tm.GetChar()
		if e != nil {
			return
		}

		tm.HandleSignals(a)

		if k == tm.UpKeyCode && curPos != 0 {
			cursor.Backward(len(choices[curPos]) + 3)
			c.print_line(choices[curPos])

			curPos = curPos - 1
			cursor.Up(1).Backward(len(choices[curPos+1]) + 3)

			c.highlight_line(choices[curPos])

		} else if k == tm.DownKeyCode && curPos < l-1 {
			cursor.Backward(len(choices[curPos]) + 3)
			c.print_line(choices[curPos])

			curPos = curPos + 1
			cursor.Down(1).Backward(len(choices[curPos-1]) + 3)

			c.highlight_line(choices[curPos])
		} else if a == tm.Enter {
			break
		}
	}

	c.Value = choices[curPos]
	cursor.Down(l - curPos)

	for l > -1 {
		cursor.Up(1)
		c.Theme.Write([]byte(tm.ClearLine))
		l = l - 1
	}
	c.Theme.Printf("%s ", label)
	c.Theme.Highlight("%s\n", c.Value)

	return
}

func (c *ListView) highlight_line(s string) {
	c.Theme.Highlight(" > %s", s)
}

func (c *ListView) print_line(s string) {
	c.Theme.Printf("   %s", s)
}

func (c *ListView) GetValue() interface{} {
	return c.Value
}

func (c *ListView) GetName() string {
	return c.Name
}

func (c *ListView) SetTheme(theme *tm.Theme) {
	c.Theme = theme
}
