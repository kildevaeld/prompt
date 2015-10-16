package form

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/oleiade/reflections"
)

func reflectField(fieldName string, target interface{}) (Field, error) {

	formTag, err := reflections.GetFieldTag(target, fieldName, "form")

	if err != nil {
		return nil, err
	}

	if formTag == "" {
		s, e := formTagFromField(fieldName, target)
		if e != nil || s == "" {
			return nil, e
		}
		formTag = s
	}

	if !validFieldType(fieldName, target, formTag) {
		return nil, fmt.Errorf("target field %s : %s", fieldName, formTag)
	}

	var message, choicesString string
	var choices []string

	message, err = reflections.GetFieldTag(target, fieldName, "message")

	if err != nil {
		return nil, err
	}

	if message == "" {
		message = fieldName
	}

	choicesString, err = reflections.GetFieldTag(target, fieldName, "choices")

	if err != nil {
		return nil, err
	}

	if choicesString != "" {
		choices = strings.Split(choicesString, ",")
		for i, c := range choices {
			choices[i] = strings.Trim(c, " ")
		}
	}

	var field Field
	switch formTag {
	case "input":
		field = &Input{
			Name:    fieldName,
			Message: message,
		}
	case "confirm":
		field = &Confirm{
			Name:    fieldName,
			Message: message,
		}
	case "password":
		field = &Password{
			Name:    fieldName,
			Message: message,
		}
	case "list":
		field = &List{
			Name:    fieldName,
			Message: message,
			Choices: choices,
		}
	case "checkbox":
		field = &Checkbox{
			Name:    fieldName,
			Message: message,
			Choices: choices,
		}
	}

	return field, nil

}

func formTagFromField(fieldName string, target interface{}) (string, error) {
	fieldKind, err := reflections.GetFieldKind(target, fieldName)
	f, _ := reflections.GetField(target, fieldName)
	fmt.Printf("Field %v\n", f)
	if err != nil {
		return "", err
	}

	choices, _ := reflections.GetFieldTag(target, fieldName, "choices")

	switch fieldKind {
	case reflect.String:
		if choices != "" {
			return "list", nil
		}
		return "input", nil
	case reflect.Slice:

		return "checkbox", nil
	case reflect.Bool:
		return "confirm", nil
	}
	return "", nil
}

func validFieldType(fieldName string, target interface{}, fieldType string) bool {
	fieldKind, err := reflections.GetFieldKind(target, fieldName)

	if err != nil {
		return false
	}

	return fieldKind == fieldKindFromFieldType(fieldType)

}

func fieldKindFromFieldType(fieldType string) reflect.Kind {
	switch fieldType {
	case "input", "list", "password":
		return reflect.String
	case "checkbox":
		return reflect.Slice
	case "confirm":
		return reflect.Bool
	}
	return reflect.Invalid
}
