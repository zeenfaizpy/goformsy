package validators

type ValidationError struct {
	fieldName string
	text      string
}

func New(fieldName string, text string) error {
	return &ValidationError{fieldName, text}
}

func (v *ValidationError) Error() string {
	return v.text
}
