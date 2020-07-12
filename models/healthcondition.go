package models

// HealthCondition used to representate health condition
type HealthCondition struct {
	Weight      float32       `json:"weight"`
	Height      float32       `json:"height"`
	Temperature float32       `json:"temperature"`
	Pressure    BloodPressure `json:"bloodPressure"`
}
