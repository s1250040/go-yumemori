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
	sampling_start_time time.Time `gorm:"type:datetime(6)"`
}

// test時
// type State struct {
// 	State int `json:"state"`
// }
