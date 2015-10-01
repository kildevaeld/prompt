package widgets

import (
	tm "github.com/kildevaeld/prompt/terminal"
)

type InputView struct {
	Theme *tm.Theme
	Name  string
	Label string
	Value string
}

func (c *InputView) Render() {

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
	var buffer []byte

	for {
		a, k, _ := tm.GetChar()
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
			c.Value = string(buffer)
			break
		} else if k == tm.RightKeyCode {
			if x < len(buffer)-1 {
				x++
				cursor.Forward(1)
			}
			continue
		} else if k == tm.LeftKeyCode {
			if x > 0 {
				x--
				cursor.Backward(1)
			}
			continue
		} else if k == tm.UpKeyCode || k == tm.DownKeyCode {
			continue
		}

		if len(buffer) == x {
			buffer = append(buffer, byte(a))
		} else {
			buffer[x] = byte(a)
		}

		c.Theme.WriteString(c.Theme.Input.Color(string(a)))

		x++
	}

	cursor.Backward(x)

	c.Theme.Highlight("%s\n", buffer)
}

func (c *InputView) GetValue() interface{} {
	return c.Value
}

func (c *InputView) GetName() string {
	return c.Name
}

func (c *InputView) SetTheme(theme *tm.Theme) {
	c.Theme = theme
}
