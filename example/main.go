package main

import (
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
