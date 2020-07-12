package models

// BloodPressure used to representate blood pressure in mmHg
type BloodPressure struct {
	Diastolic int16 `json:"diastolic"`
	Systolic  int16 `json:"Systolic"`
}
