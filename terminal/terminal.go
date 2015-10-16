package terminal

import (
	"os"
	"syscall"

	acsii "github.com/kildevaeld/go-acsii"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/pkg/term"
)

const (
	HideCursor = "\033[?25l"
	ShowCursor = "\033[?25h"
	//Gray        = "\033[90m"
	ClearLine    = "\r\033[0K"
	UpKeyCode    = 38
	DownKeyCode  = 40
	RightKeyCode = 39
	LeftKeyCode  = 37
	Enter        = 13
	Backspace    = 127
	Space        = 32
)

const (
	keyCtrlC     = 3
	keyCtrlD     = 4
	keyCtrlU     = 21
	keyCtrlZ     = 26
	keyEnter     = '\r'
	keyEscape    = 27
	keyBackspace = 127
	keyUnknown   = 0xd800 /* UTF-16 surrogate area */ + iota
	keyUp
	keyDown
	keyLeft
	keyRight
	keyAltLeft
	keyAltRight
	keyHome
	keyEnd
	keyDeleteWord
	keyDeleteLine
	keyClearScreen
	keyPasteStart
	keyPasteEnd
)

// Returns either an ascii code, or (if input is an arrow) a Javascript key code.
func GetChar() (ascii int, keyCode int, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	}
	//fmt.Printf("%v", bytes)
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
		// Three-character control sequence, beginning with "ESC-[".

		// Since there are no ASCII codes for arrow keys, we use
		// Javascript key codes.
		if bytes[2] == 65 {
			// Up
			keyCode = 38
		} else if bytes[2] == 66 {
			// Down
			keyCode = 40
		} else if bytes[2] == 67 {
			// Right
			keyCode = 39
		} else if bytes[2] == 68 {
			// Left
			keyCode = 37
		}
		ascii = int(bytes[2])
	} else if numRead == 1 {
		ascii = int(bytes[0])
	} else {
		// Two characters read??
	}
	t.Restore()
	t.Close()
	return
}

func GetSize() (int, int, error) {
	w, h, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return -1, -1, err
	}
	return w, h, nil
}

func Save() {
	os.Stdout.Write([]byte(acsii.ESC + "?1049h\033[H"))
}

func Restore() {
	os.Stdout.Write([]byte(acsii.ESC + "?1049l"))
}

func Clear() {
	os.Stdout.Write([]byte(acsii.ESC + "2J"))
}

func HandleSignals(c int) {
	pid := syscall.Getpid()
	cur := Cursor{os.Stdout}
	switch c {
	case keyCtrlC:
		cur.Show()
		syscall.Kill(pid, syscall.SIGINT)
	case keyCtrlZ:
		cur.Show()
		syscall.Kill(pid, syscall.SIGTSTP)
	}
}
