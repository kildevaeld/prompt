package main

import (
	"fmt"
	"time"

	"github.com/kildevaeld/prompt"
	"github.com/kildevaeld/prompt/terminal"
	"github.com/kildevaeld/prompt/widgets"
)

type Result struct {
	Name     string
	Password string
	List     string
}

func main() {

	ui := prompt.NewUI()
	ui.Theme = terminal.DefaultTheme
	ui.Save() // Clear the terminal
	// or ui.Save()

	var result Result

	prompt.NewProcess("Test mig", func() error {
		time.Sleep(1 * time.Second)
		return nil
	})

	prompt.NewProgress("msg", func(fn func(string)) error {

		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			fn(fmt.Sprintf("%d/10", i))
		}

		return nil
	})

	ui.Form([]widgets.Field{
		&widgets.InputView{
			Name:  "name",
			Label: "Please enter name?",
		},
		&widgets.PasswordView{
			Name:  "password",
			Label: "Password",
		},
		&widgets.ListView{
			Name:    "List",
			Choices: []string{"Cheese", "Ham"},
		},
	}, &result)

	// ui.Restore() to restore from "Save"
	ui.Printf("%#v", result)

}
