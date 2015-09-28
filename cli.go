package cli

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/pkg/term"
	"github.com/ttacon/chalk"
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
func getChar() (ascii int, keyCode int, err error) {
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

func MoveCursor(x, y int) {
	os.Stdout.Write([]byte(fmt.Sprintf("\033[%d;%dH", x, y)))
}

func moveUp(i int) {
	os.Stdout.Write([]byte(fmt.Sprintf("\033[%dA", i)))
}

func moveDown(i int) {
	os.Stdout.Write([]byte(fmt.Sprintf("\033[%dB", i)))
}

func moveForward(i int) {
	os.Stdout.Write([]byte(fmt.Sprintf("\033[%dC", i)))
}

func moveBack(i int) {
	os.Stdout.Write([]byte(fmt.Sprintf("\033[%dD", i)))
}

func Save() {
	os.Stdout.Write([]byte("\033[?1049h\033[H"))
}

func Restore() {
	os.Stdout.Write([]byte("\033[?1049l"))
}

func Clear() {
	os.Stdout.Write([]byte("\033[2J"))
}

func Confirm(msg string) bool {
	var ans string
	fmt.Printf(fmt.Sprintf("%s%s%s [yn]? %s", ClearLine, Gray, msg, chalk.Reset))

	a, k, _ := getChar()
	fmt.Printf("%s %s", string(a), string(k))
	ans = string(a)
	if ans == "y" || ans == "Y" {
		return true
	} else if ans == "n" || ans == "n" {
		fmt.Print("\n")
		return false
	} else {
		fmt.Printf("\r%s%s %splease enter %s(es) or n(o)", Gray, msg, chalk.Cyan.NewStyle(), chalk.Bold.TextStyle("y"))
		time.Sleep(1 * time.Second)
		return Confirm(msg)
	}
}

func handleSignals(c int) {
	pid := syscall.Getpid()
	switch c {
	case keyCtrlC:
		syscall.Kill(pid, syscall.SIGINT)
	case keyCtrlZ:
		syscall.Kill(pid, syscall.SIGTSTP)
	}
}

func Password(msg string) string {
	writer := os.Stdout
	writer.Write([]byte(msg + " "))
	x := 0
	buffer := ""
	for {
		a, _, _ := getChar()
		handleSignals(a)
		if a == keyCtrlC {
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			return ""
		} else if a == Backspace {
			if x == 0 {
				continue
			}

			writer.Write([]byte("\b \b"))

			x--
			buffer = buffer[0:x]
			continue
		} else if a == keyEnter {
			moveDown(1)
			writer.Write([]byte("\r"))
			return buffer
		}

		buffer += string(a)

		writer.Write([]byte("*"))
		//moveForward(0)
		x++
	}
}
