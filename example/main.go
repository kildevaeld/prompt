package main

import "github.com/kildevaeld/prompt"

type Result struct {
	Name     string
	Password string
	List     string
}

func main() {

	ui := prompt.NewUI()
	ui.Clear()
	ui.Theme.Indent = "  "

	/*var result Result
	ui.Form([]prompt.Field{
		&prompt.InputView{
			Name:  "name",
			Label: "Okiedokie",
		},
		&prompt.PasswordView{
			Name:  "password",
			Label: "Password",
		},
		&prompt.ListView{
			Name: "list",
			Choices: []string{
				"Ost", "Makral",
			},
		},
	}, &result)*/

	ui.List("List", []string{"Første valg", "Andet valg"})

	//fmt.Printf("%#v", result)
	//ui.Move(10, 20)
	//ui.Password("Password:")
	//ui.Confirm("No?")

}
