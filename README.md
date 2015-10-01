# go-prompt
Golang terminal

## Usage

```go

import (
	"fmt"
	"github.com/kildevaeld/prompt"
	"github.com/kildevaeld/prompt/widgets"
)

type Result struct {
	Name     string
	Password string
	List 	 string
}

ui := prompt.NewUI()

prompt.Clear() // Clear the terminal
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
		Name: "List",
		Choices: []string{"Cheese", "Ham"},
	},
}, &result)

// ui.Restore()
fmt.Printf("%#v", result)

// outputs: Result{Name:"John Doe", Password:"password", List:"Ham"}

```
