package models

// Location used to representate location
type Location struct {
	Coordinate Coordinate `json:"coordinate"`
	Name       string     `json:"name"`
	Contact    Contact    `json:"contact"`
	Owner      string     `json:"owner"`
}
