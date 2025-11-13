package forms

const BaseErrorKey = "base"

type FormInterface interface {
	HasErrors() bool
	Errors() map[string]Errors
	HasBaseErrors() bool
	BaseErrors() Errors
	AddBaseError(string)
	Fields() map[string]FieldInterface
}

type Form struct {
	baseErrorSl []string
	fieldMap    map[string]FieldInterface
}

func NewForm() *Form {
	return &Form{
		baseErrorSl: []string{},
		fieldMap:    map[string]FieldInterface{},
	}
}

func (f *Form) AddField(name string, fi FieldInterface) {
	f.fieldMap[name] = fi
}

func (f *Form) Fields() map[string]FieldInterface {
	return f.fieldMap
}

func (f *Form) HasBaseErrors() bool {
	return len(f.baseErrorSl) > 0
}

func (f *Form) BaseErrors() []string {
	return f.baseErrorSl
}

func (f *Form) AddBaseError(errStr string) {
	f.baseErrorSl = append(f.baseErrorSl, errStr)
}

func (f *Form) Errors() map[string][]string {
	errs := map[string][]string{}
	if f.HasBaseErrors() {
		errs["base"] = f.BaseErrors()
	}

	for k, v := range f.Fields() {
		if v.HasErrors() {
			errs[k] = v.Errors()
		}
	}
	return errs
}

func (f *Form) HasErrors() bool {
	return len(f.Errors()) > 0
}
