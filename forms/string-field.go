package forms

type StringField struct {
	*Field[string]
}

func NewStringField(f *Form, name string, val string) *StringField {
	sfield := &StringField{
		Field: newField[string](name),
	}
	sfield.Set(val)
	f.AddField(name, sfield)
	return sfield
}

func (field *StringField) Set(v string) {
	if v == "" {
		field.Clear()
		return
	}

	field.present = true
	field.formVal = v
	field.val = v
}
