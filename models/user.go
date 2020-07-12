package models

// User used to representate user
type User struct {
	Fullname   string `json:"fullname"`
	Locked     bool   `json:"locked"`
	Registered bool   `json:"registered"`
	Deleted    bool   `json:"deleted"`
	Owner      string `json:"owner"`
}
