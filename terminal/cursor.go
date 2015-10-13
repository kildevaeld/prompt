package terminal

import (
	"io"

	acsii "github.com/kildevaeld/go-acsii"
)

type Cursor struct {
	Writer io.Writer
}

func (c Cursor) Move(x, y int) Cursor {
	return c.writeString(acsii.CursorMove(x, y))

}

func (c Cursor) Forward(x int) Cursor {
	return c.writeString(acsii.CursorForward(x))
}

func (c Cursor) Backward(x int) Cursor {
	return c.writeString(acsii.CursorBackward(x))
}

func (c Cursor) Up(y int) Cursor {
	return c.writeString(acsii.CursorUp(y))
}

func (c Cursor) Down(y int) Cursor {
	return c.writeString(acsii.CursorDown(y))
}

func (c Cursor) Hide() Cursor {
	return c.writeString(acsii.CursorHide)
}

func (c Cursor) Show() Cursor {
	return c.writeString(acsii.CursorShow)
}

func (c Cursor) writeString(str string) Cursor {
	c.Writer.Write([]byte(str))
	return c
}
