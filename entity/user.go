package entity

// User is user models property
type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	State int    `json:"state"`
}
