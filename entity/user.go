package entity

// User is user models property
type User struct {
	ID uint `json:"id"`
	// FirstName string `json:"firstname"`
	// LastName  string `json:"lastname"`
	Name   string `json:"name"`
	States []State
}

type State struct {
	State int `json:"state"`
}
