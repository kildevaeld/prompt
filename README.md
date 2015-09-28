# go-prompt
Golang terminal

## Usage
 
```go

type Result struct {
	Name     string
	Password string
}

ui := cli.NewUI()

ui.Clear() // Clear the terminal

var result Result

ui.Form([]cli.Field{
	&cli.InputView{
		Name:  "name",
		Label: "Please enter name?",
	},
	&cli.PasswordView{
		Name:  "password",
		Label: "Password",
	},
}, &result)

fmt.Printf("%#v", result)

// outputs: Result{Name:"John Doe", Password:"password"}

```
