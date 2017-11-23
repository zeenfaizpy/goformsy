package forms

import (
	"github.com/zeenfaizpy/goformsy/fields"
)

type FormFields map[string]fields.Field

type Form struct {
	Fields      FormFields
	FieldsOrder []string
	Data        map[string]interface{}
	CleanedData map[string]interface{}
	Errors      map[string][]error
	Initial     map[string]string
}

func (f *Form) IsBound() bool {
	if len(f.Data) > 0 {
		return true
	}
	return false
}

func (f *Form) HasErrors() bool {
	if len(f.Errors) > 0 {
		return true
	}
	return false
}

func (f *Form) IsValid() bool {
	if f.IsBound() && f.HasErrors() {
		return true
	}
	return false
}

func (f *Form) Clean() map[string]interface{} {
	return f.CleanedData
}

func (f *Form) FullClean() bool {

	f.CleanedData = make(map[string]interface{})
	f.Errors = make(map[string][]error)

	for _, fieldName := range f.FieldsOrder {
		if field, ok := f.Fields[fieldName]; ok {
			value, err = field.Clean(f.Data)
			if err != nil {
				f.CleanedData[fieldName] = value
			} else {
				append(f.Errors[fieldName], err)
			}
		}
	}
}
