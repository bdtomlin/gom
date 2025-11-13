package forms

type IntField struct {
	*Field[int]
}

func NewIntField(f *Form, name string, val string) *IntField {
	ifield := &IntField{
		Field: newField[int](name),
	}
	ifield.Set(val)
	f.AddField(name, ifield)
	return ifield
}

func (field *IntField) Val() int {
	return field.val
}

func (field *IntField) Set(numStr string) {
	if numStr == "" {
		field.Clear()
		return
	}

	field.present = true

	str, num, err := ParseInt(numStr)
	if err != nil {
		field.AddError("invalid integer")
	}
	field.formVal = str
	field.val = num
}
