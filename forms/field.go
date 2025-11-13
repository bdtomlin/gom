package forms

type Errors []string

type FieldInterface interface {
	Name() string
	FormVal() string
	HasErrors() bool
	Errors() Errors
	Set(string)
}

type FieldType interface {
	~string | ~int | ~float64
}

type Field[T FieldType] struct {
	name       string
	val        T
	present    bool
	formVal    string
	errorSlice []string
}

func newField[T FieldType](name string) *Field[T] {
	return &Field[T]{
		name:       name,
		errorSlice: []string{},
	}
}

func (fv *Field[T]) Clear() {
	var v T

	fv.val = v
	fv.formVal = ""
	fv.present = false
}

func (fld *Field[T]) Name() string {
	return fld.name
}

func (fld *Field[T]) HasErrors() bool {
	return len(fld.errorSlice) > 0
}
func (fld *Field[T]) Errors() Errors {
	return fld.errorSlice
}

func (fld *Field[T]) AddError(errStr string) {
	if fld.errorSlice == nil {
		fld.errorSlice = []string{}
	}
	fld.errorSlice = append(fld.errorSlice, errStr)
}

func (field *Field[T]) Present() bool {
	return field.present
}

func (field *Field[T]) FormVal() string {
	return field.formVal
}

func (field *Field[T]) Val() T {
	return field.val
}
