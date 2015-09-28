package cli

import (
	"fmt"
	"io"
	"os"
)

type Theme struct {
	Background, Foreground, HighlightForeground,
	HighlightBackground, Input Color
	Indent string
	writer io.Writer
}

func (t *Theme) Printf(msg string, args ...interface{}) {
	t.WriteString(t.Foreground.Color(fmt.Sprintf(msg, args...)))
}

func (t *Theme) Highlight(msg string, args ...interface{}) {
	t.WriteString(t.HighlightForeground.Color(fmt.Sprintf(msg, args...)))
}

func (t *Theme) WriteString(msg string) {
	t.Write([]byte(msg))
}

func (t *Theme) Write(bytes []byte) (int, error) {
	return t.writer.Write(bytes)
}

var DefaultTheme = &Theme{
	Background:          Black,
	Foreground:          Gray,
	HighlightForeground: Cyan,
	HighlightBackground: Gray,
	Input:               White,
	writer:              os.Stdout,
}
