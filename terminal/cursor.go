package terminal

import (
	"io"

	ascii "github.com/kildevaeld/go-ascii"
)

type Cursor struct {
	Writer io.Writer
}

func (c Cursor) Move(x, y int) Cursor {
	return c.writeString(ascii.CursorMove(x, y))

}

func (c Cursor) Forward(x int) Cursor {
	return c.writeString(ascii.CursorForward(x))
}

func (c Cursor) Backward(x int) Cursor {
	return c.writeString(ascii.CursorBackward(x))
}

func (c Cursor) Up(y int) Cursor {
	return c.writeString(ascii.CursorUp(y))
}

func (c Cursor) Down(y int) Cursor {
	return c.writeString(ascii.CursorDown(y))
}

func (c Cursor) Hide() Cursor {
	return c.writeString(ascii.CursorHide)
}

func (c Cursor) Show() Cursor {
	return c.writeString(ascii.CursorShow)
}

func (c Cursor) writeString(str string) Cursor {
	c.Writer.Write([]byte(str))
	return c
}
