package cli

type InputView struct {
	theme *Theme
	Name  string
	Label string
	Value string
}

func (c *InputView) Render() {

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
	var buffer []byte

	for {
		a, k, _ := getChar()
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
			c.Value = string(buffer)
			break
		} else if k == RightKeyCode {
			if x < len(buffer)-1 {
				x++
				cursor.Forward(1)
			}
			continue
		} else if k == LeftKeyCode {
			if x > 0 {
				x--
				cursor.Backward(1)
			}
			continue
		} else if k == UpKeyCode || k == DownKeyCode {
			continue
		}

		if len(buffer) == x {
			buffer = append(buffer, byte(a))
		} else {
			buffer[x] = byte(a)
		}

		c.theme.WriteString(c.theme.Input.Color(string(a)))

		x++
	}

	cursor.Backward(x)

	c.theme.Highlight("%s\n", buffer)
}

func (c *InputView) GetValue() interface{} {
	return c.Value
}

func (c *InputView) GetName() string {
	return c.Name
}

func (c *InputView) SetTheme(theme *Theme) {
	c.theme = theme
}
