package widgets

import (
	"errors"

	"github.com/kildevaeld/prompt/terminal"
	"github.com/mitchellh/mapstructure"
)

type Field interface {
	Render()
	GetValue() interface{}
	GetName() string
	SetTheme(theme *terminal.Theme)
}

type Form struct {
	fields []Field
	Theme  *terminal.Theme
	Value  map[string]interface{}
}

func (f *Form) Render() {
	values := make(map[string]interface{})

	for _, field := range f.fields {
		field.SetTheme(f.Theme)
		field.Render()
		values[field.GetName()] = field.GetValue()
	}
	f.Value = values
}

func (f *Form) GetValue(v interface{}) error {
	if f.Value == nil {
		return errors.New("no value")
	}
	return mapstructure.Decode(f.Value, v)
}

func NewForm(theme *terminal.Theme, fields []Field) *Form {
	return &Form{fields, theme, nil}
}
