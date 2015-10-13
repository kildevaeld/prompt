package terminal

import (
	"fmt"
	"io"
	"os"
)

type Theme struct {
	Background, Foreground, HighlightForeground,
	HighlightBackground, Input, Error, Success Color
	Indent string
	Writer io.Writer
	Cursor Cursor
}

func (t *Theme) Printf(msg string, args ...interface{}) int {
	return t.WriteString(t.Foreground.Color(fmt.Sprintf(msg, args...)))
}

func (t *Theme) Highlight(msg string, args ...interface{}) int {
	return t.WriteString(t.HighlightForeground.Color(fmt.Sprintf(msg, args...)))
}

func (t *Theme) WriteString(msg string) int {
	l := len(msg)
	t.Write([]byte(msg))
	return l
}

func (t *Theme) Write(bytes []byte) (int, error) {
	return t.Writer.Write(bytes)
}

var DefaultTheme = &Theme{
	Background:          Black,
	Foreground:          Gray,
	HighlightForeground: Cyan,
	HighlightBackground: Gray,
	Error:               Red,
	Success:             Green,
	Input:               White,
	Writer:              os.Stdout,
	Cursor:              Cursor{os.Stdout},
}
