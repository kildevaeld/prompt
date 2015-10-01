package prompt

type ListView struct {
	Name    string
	Label   string
	Value   string
	Choices []string
	theme   *Theme
}

func (c *ListView) Render() {
	choices := c.Choices
	if c.theme == nil {
		c.theme = DefaultTheme
	}

	cursor := Cursor{
		writer: c.theme,
	}

	label := c.Label
	if label == "" {
		label = c.Name
	}
	cursor.Hide()
	c.theme.Printf("%s\n", label)

	for i, s := range choices {
		if i == len(choices)-1 {
			c.highlight_line(s)
		} else {
			c.print_line(s)
		}
		c.theme.WriteString("\n")

	}
	l := len(choices)

	cursor.Up(1)
	curPos := l - 1
	for {
		a, k, e := getChar()
		if e != nil {
			return
		}

		handleSignals(a)

		if k == UpKeyCode && curPos != 0 {
			cursor.Backward(len(choices[curPos]) + 3)
			c.print_line(choices[curPos])

			curPos = curPos - 1
			cursor.Up(1).Backward(len(choices[curPos+1]) + 3)

			c.highlight_line(choices[curPos])

		} else if k == DownKeyCode && curPos < l-1 {
			cursor.Backward(len(choices[curPos]) + 3)
			c.print_line(choices[curPos])

			curPos = curPos + 1
			cursor.Down(1).Backward(len(choices[curPos-1]) + 3)

			c.highlight_line(choices[curPos])
		} else if a == Enter {
			break
		}
	}

	c.Value = choices[curPos]
	cursor.Down(l - curPos)

	for l > -1 {
		cursor.Up(1)
		c.theme.Write([]byte(ClearLine))
		l = l - 1
	}
	c.theme.Printf("%s ", label)
	c.theme.Highlight("%s\n", c.Value)

	return
}

func (c *ListView) highlight_line(s string) {
	c.theme.Highlight(" > %s", s)
}

func (c *ListView) print_line(s string) {
	c.theme.Printf("   %s", s)
}

func (c *ListView) GetValue() interface{} {
	return c.Value
}

func (c *ListView) GetName() string {
	return c.Name
}

func (c *ListView) SetTheme(theme *Theme) {
	c.theme = theme
}
