# go-prompt
Golang terminal

## Usage
 
```go

type Result struct {
	Name     string
	Password string
	List 	 string
}

ui := cli.NewUI()

ui.Clear() // Clear the terminal
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
