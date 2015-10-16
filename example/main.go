package main

import (
	"fmt"
	"log"

	"github.com/kildevaeld/prompt"
	"github.com/kildevaeld/prompt/terminal"
)

type Result struct {
	Name     string   `message:"Enter name please?"`
	Password string   `form:"password"`
	List     string   `form:"list" choices:"Choice 1, Choice 2"`
	Checkbox []string `form:"checkbox" choices:"Choice 1, Choice 2"`
}

func main() {

	ui := prompt.NewUI()
	ui.Theme = terminal.DefaultTheme
	//ui.Clear() // Clear the terminal
	// or ui.Save()

	//var result Result

	/*ui.FormWithFields([]form.Field{
		&form.Input{
			Name:    "name",
			Message: "Please enter name?",
		},
		&form.Password{
			Name:    "password",
			Message: "Password",
		},
		&form.List{
			Name:    "List",
			Choices: []string{"Cheese", "Ham"},
		},
	}, &result)*/

	//fmt.Printf("%#v\n", result)
	// Or
	var ret Result
	if err := ui.Form(&ret); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", ret)
	//form.Run()
	// ui.Restore() to restore from "Save"
	//ui.Printf("%#v", result)

}
