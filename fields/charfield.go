package fields

import (
	"fmt"
	"github.com/zeenfaizpy/goformsy/utils"
	"github.com/zeenfaizpy/goformsy/widgets"
    "github.com/zeenfaizpy/goformsy/validators"
)

type CharField struct {
	BaseField
	Initial   string
	Widget    widgets.Widget
	MaxLength int
	MinLength int
	Strip     bool
}

func (f *CharField) New(options Options) CharField {
	field := CharField{}

	for optionName, optionValue := range options {
		switch optionName {
		case "Required":
			if v, ok := optionValue.(bool); ok {
				field.Required = v
			}
		case "Label":
			if v, ok := optionValue.(string); ok {
				field.Label = v
			}
		case "HelpText":
			if v, ok := optionValue.(string); ok {
				field.HelpText = v
			}
		case "Initial":
			if v, ok := optionValue.(string); ok {
				field.Initial = v
			}
		case "MaxLength":
			if v, ok := optionValue.(int); ok {
				field.MaxLength = v
			}
		case "MinLength":
			if v, ok := optionValue.(int); ok {
				field.MinLength = v
			}
		case "Strip":
			if v, ok := optionValue.(bool); ok {
				field.Strip = v
			}
		}
	}

	return field
}

func (f *CharField) getValue(data map[string]interface{}) (string, error) {
	if value, ok := data[f.Name]; ok {
		if v, ok := value.(string); ok {
			var err = validators.New(f.Name, "String only allowed")
			return v, err
		}
	}
	
	return "", nil
}


func (f *CharField) Validate(value string) (string, error) {
	var emptyValues = []string{"", " "}
	if f.Required {
		if utils.Contains(emptyValues, value) {
            var err = validators.New(f.Name, "This field is Required")
			return value, err
		}
	}

	if len(value) < f.MinLength {
		var err = validators.New(f.Name, fmt.Sprintf("Should have atleast %d characters", f.MinLength))
		return value, err
	}

	if len(value) > f.MaxLength {
		var err = validators.New(f.Name, fmt.Sprintf("Only %d characters allowed", f.MaxLength))
		return value, err
	}

	return value, nil
}

func (f *CharField) Clean(data map[string]interface{}) (string, error) {
	value, err := f.getValue(data)
	v, err := f.Validate(value)
	return v, err
}
