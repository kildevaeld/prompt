package main

import (
	"fmt"

	"github.com/kildevaeld/cli"
)

type Result struct {
	Name     string
	Password string
}

func main() {

	ui := cli.NewUI()
	ui.Clear()
	ui.Theme.Indent = "  "

	var result Result
	ui.Form([]cli.Field{
		&cli.InputView{
			Name:  "name",
			Label: "Okiedokie",
		},
		&cli.PasswordView{
			Name:  "password",
			Label: "Password",
		},
		&cli.ListView{
			Name: "list",
			Choices: []string{
				"Ost", "Makral",
			},
		},
	}, &result)

	fmt.Printf("%#v", result)
	//ui.Move(10, 20)
	//ui.Password("Password:")
	//ui.Confirm("No?")

}
