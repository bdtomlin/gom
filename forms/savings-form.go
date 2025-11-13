package forms

import (
	"net/http"
)

type SavingsForm struct {
	*Form
	NumOffices              *IntField
	Volume                  *IntField
	AvgStatements           *IntField
	AvgPerStatement         *IntField
	TimeBilling             *IntField
	TimeCollecting          *IntField
	TimeReconciling         *IntField
	TimeProcessing          *IntField
	FirstName               *StringField
	LastName                *StringField
	Email                   *StringField
	Phone                   *StringField
	NumberOfYears           float64
	HourlyPay               float64
	HourlyPayCollecting     float64
	PercentSavedBilling     float64
	PercentSavedCollecting  float64
	PercentSavedReconciling float64
	PercentSavedProcessing  float64
	PercentExtraCollected   float64
}

func NewSavingsForm() *SavingsForm {
	f := NewForm()
	return &SavingsForm{
		Form:                    f,
		NumOffices:              NewIntField(f, "NumOffices", ""),
		Volume:                  NewIntField(f, "Volume", ""),
		AvgStatements:           NewIntField(f, "AvgStatements", ""),
		AvgPerStatement:         NewIntField(f, "AvgPerStatement", ""),
		TimeBilling:             NewIntField(f, "TimeBilling", ""),
		TimeCollecting:          NewIntField(f, "TimeCollecting", ""),
		TimeReconciling:         NewIntField(f, "TimeReconciling", ""),
		TimeProcessing:          NewIntField(f, "TimeProcessing", ""),
		FirstName:               NewStringField(f, "FirstName", ""),
		LastName:                NewStringField(f, "LastName", ""),
		Email:                   NewStringField(f, "Email", ""),
		Phone:                   NewStringField(f, "Phone", ""),
		NumberOfYears:           3.0,
		HourlyPay:               44.12,
		HourlyPayCollecting:     22.82,
		PercentSavedBilling:     0.2,
		PercentSavedCollecting:  0.3,
		PercentSavedReconciling: 0.25,
		PercentSavedProcessing:  0.9,
		PercentExtraCollected:   0.0974,
	}
}

func DefaultSavingsForm() *SavingsForm {
	f := NewSavingsForm()
	f.NumOffices.Set("1")
	f.Volume.Set("25000")
	f.AvgStatements.Set("220")
	f.AvgPerStatement.Set("40")
	f.TimeBilling.Set("20")
	f.TimeCollecting.Set("10")
	f.TimeReconciling.Set("10")
	f.TimeProcessing.Set("1")

	return f
}

func (sf *SavingsForm) Decode(r *http.Request) {
	for name, field := range sf.Fields() {
		field.Set(r.PostFormValue(name))
	}
}

func (sf *SavingsForm) Validate() {
	sf.ValidateNum(sf.NumOffices, 0)
	sf.ValidateNum(sf.Volume, 0)
	sf.ValidateNum(sf.AvgStatements, 0)
	sf.ValidateNum(sf.AvgPerStatement, 0)
	sf.ValidateNum(sf.TimeBilling, 0)
	sf.ValidateNum(sf.TimeCollecting, 0)
	sf.ValidateNum(sf.TimeReconciling, 0)
	sf.ValidateNum(sf.TimeProcessing, 0)
}

func (sf *SavingsForm) ValidateNum(field *IntField, gt int) {
	if !field.Present() {
		field.AddError("must be present")
		return
	}
	if field.Val() <= 0 {
		field.AddError("must be greater than 0")
	}
}

func (sf *SavingsForm) BillingSavings() int {
	prod :=
		float64(sf.NumOffices.Val()) *
			float64(sf.TimeBilling.Val()) *
			52.0 *
			sf.HourlyPay *
			sf.PercentSavedBilling *
			sf.NumberOfYears

	return FloatToInt(prod)
}

func (sf *SavingsForm) CollectionSavings() int {
	prod :=
		float64(sf.NumOffices.Val()) *
			float64(sf.TimeCollecting.Val()) *
			52.0 *
			sf.HourlyPayCollecting *
			sf.PercentSavedCollecting *
			sf.NumberOfYears

	return FloatToInt(prod)
}

func (sf *SavingsForm) ReconciliationSavings() int {
	prod :=
		float64(sf.NumOffices.Val()) *
			float64(sf.TimeReconciling.Val()) *
			52.0 *
			sf.HourlyPay *
			sf.PercentSavedReconciling *
			sf.NumberOfYears

	return FloatToInt(prod)
}

func (sf *SavingsForm) ProcessingSavings() int {
	prod :=
		float64(sf.NumOffices.Val()) *
			float64(sf.TimeProcessing.Val()) *
			52.0 *
			sf.HourlyPay *
			sf.PercentSavedProcessing *
			sf.NumberOfYears

	return FloatToInt(prod)
}

func (sf *SavingsForm) ExtraCollected() int {
	prod :=
		float64(sf.AvgStatements.Val()) *
			0.0966666 *
			float64(sf.AvgPerStatement.Val()) *
			sf.NumberOfYears * 12

	return FloatToInt(prod)
}
func (sf *SavingsForm) ImproveOpSavings() int {
	sum := sf.BillingSavings() +
		sf.CollectionSavings() +
		sf.ReconciliationSavings() +
		sf.ProcessingSavings()

	return sum
}

func (sf *SavingsForm) TotalValue() int {
	return sf.ExtraCollected() + sf.ImproveOpSavings()
}

func (sf *SavingsForm) MonthlyValue() int {
	return sf.TotalValue() / int(sf.NumberOfYears) / 12
}
