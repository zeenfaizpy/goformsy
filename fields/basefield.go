package fields

type Options map[string]interface{}


type Field interface {
	New(Options) Field
}

type BaseField struct {
	Name string
	Label    string
	HelpText string
	Required bool
}
