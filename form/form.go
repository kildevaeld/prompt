package form

import (
	"errors"

	"github.com/kildevaeld/prompt/terminal"
	"github.com/mitchellh/mapstructure"
	"github.com/oleiade/reflections"
)

type Field interface {
	Run()
	GetValue() interface{}
	GetName() string
}

type Form struct {
	fields []Field
	Theme  *terminal.Theme
	Value  map[string]interface{}
}

func (f *Form) Run() {
	values := make(map[string]interface{})

	for _, field := range f.fields {
		//field.SetTheme(f.Theme)
		field.Run()
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

func FormFromStruct(theme *terminal.Theme, target interface{}) error {
	fields, err := reflections.Fields(target)

	if err != nil {
		return err
	}

	var formFields []Field

	for _, fieldName := range fields {

		field, err := reflectField(fieldName, target)
		if err != nil {
			return err
		}
		formFields = append(formFields, field)

	}

	form := NewForm(theme, formFields)
	form.Run()

	return form.GetValue(target)

}
