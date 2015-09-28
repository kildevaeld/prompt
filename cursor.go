package cli

import (
	"fmt"
	"io"
)

type Cursor struct {
	writer io.Writer
}

func (c Cursor) move(i int, direction rune) Cursor {
	c.writer.Write([]byte(fmt.Sprintf("\033[%d%s", i, string(direction))))
	return c
}

func (c Cursor) Move(x, y int) Cursor {
	c.writer.Write([]byte(fmt.Sprintf("\033[%d;%dH", x, y)))
	return c
}

func (c Cursor) Forward(x int) Cursor {
	return c.move(x, 'C')
}

func (c Cursor) Backward(x int) Cursor {
	return c.move(x, 'D')
}

func (c Cursor) Up(y int) Cursor {
	return c.move(y, 'A')
}

func (c Cursor) Down(y int) Cursor {
	return c.move(y, 'B')
}

func (c Cursor) Hide() Cursor {
	c.writer.Write([]byte("\033[?25l"))
	return c
}

func (c Cursor) Show() Cursor {
	c.writer.Write([]byte("\033[?25h"))
	return c
}
