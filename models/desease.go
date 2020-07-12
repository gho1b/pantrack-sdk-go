package models

// Desease used to representate desease
type Desease struct {
	ICD   string `json:"icd"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
