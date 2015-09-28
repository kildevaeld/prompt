package cli

type PasswordView struct {
	theme *Theme
	Name  string
	Label string
	Value string
}

func (c *PasswordView) Render() {

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

	c.theme.Printf("%s ", label)
	x := 0

	buffer := ""

	for {
		a, _, _ := getChar()
		handleSignals(a)
		if a == Backspace {
			if x == 0 {
				continue
			}
			c.theme.Write([]byte("\b \b"))

			x--
			buffer = buffer[0:x]
			continue

		} else if a == Enter {
			c.Value = buffer
			break
		}

		buffer += string(a)
		c.theme.Write([]byte(c.theme.Input.Color("*")))

		x++
	}

	cursor.Backward(x)
	str := ""
	for x > 0 {
		str += "*"
		x--
	}

	c.theme.Highlight("%s\n", str)
}

func (c *PasswordView) GetValue() interface{} {
	return c.Value
}

func (c *PasswordView) GetName() string {
	return c.Name
}

func (c PasswordView) SetTheme(theme *Theme) {
	c.theme = theme
}
