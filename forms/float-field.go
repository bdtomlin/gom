package forms

type FloatField struct {
	*Field[float64]
}

func NewFloatField(f *Form, name string, val string) *FloatField {
	ffield := &FloatField{
		Field: newField[float64](name),
	}
	ffield.Set(val)
	f.AddField(name, ffield)
	return ffield
}

func (field *FloatField) Val() float64 {
	return field.val
}

func (field *FloatField) Set(numStr string) {
	if numStr == "" {
		field.Clear()
		return
	}

	field.present = true

	s, f, err := ParseFloat(numStr)

	if err != nil {
		field.AddError("invalid decimal")
	}

	field.formVal = s
	field.val = f
}
