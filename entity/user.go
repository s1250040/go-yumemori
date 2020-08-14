package entity

import "time"

// User is user models property
type User struct {
	// ID uint `json:"id"` test時
	// FirstName string `json:"firstname"`
	// LastName  string `json:"lastname"`
	// Name string `json:"name"`　test時
	// States []State　test時
	Date string `json:"date"`
}

type Result struct {
	SamplingStartTime time.Time `sql:"not null;type:date"`
	// FK_bsb_no int
}

// test時
// type State struct {
// 	State int `json:"state"`
// }
