package cli

import (
	"fmt"
	"os"

	tm "github.com/buger/goterm"
)

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

	c.theme.Printf("%s\n", label)

	for _, s := range choices {
		c.theme.Printf("  %s\n", s)
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
			cursor.Up(1)
			curPos = curPos - 1
		} else if k == DownKeyCode && curPos < l-1 {
			cursor.Down(1)
			curPos = curPos + 1
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
	//c.theme.Write([]byte(label + " " + choices[curPos] + "\n"))
	return
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

func List(msg string, choices []string) error {
	tm.Clear()
	writer := os.Stdout
	for i, s := range choices {
		writer.Write([]byte(fmt.Sprintf("%d - %s\n", i, s)))
	}
	l := len(choices)
	writer.Write([]byte(msg))
	moveUp(1)
	curPos := l - 1
	for {
		a, k, e := getChar()
		if e != nil {
			return e
		}

		if k == UpKeyCode && curPos != 0 {
			moveUp(1)
			curPos = curPos - 1
		} else if k == DownKeyCode && curPos < l-1 {
			moveDown(1)
			curPos = curPos + 1
		} else if a == 3 {
			moveDown(l - curPos)
			writer.Write([]byte("\n"))
			return nil
		} else if a == Enter {
			break
		}
	}
	moveDown(l - curPos)

	for l > 0 {
		moveUp(1)
		writer.Write([]byte(ClearLine))
		l = l - 1
	}
	writer.Write([]byte(msg + " " + choices[curPos] + "\n"))
	return nil
}
