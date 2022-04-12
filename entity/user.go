package entity

import "time"

type User struct {
	UserId 		uint64 `gorm:"primaryKey;column:userId" json:"userId"`
	// ID 			uint64 `gorm:"primary_key;auto_increment" json:"userId"`
	Account 	string `json:"account"`
	Password  	string `json:"password"`
	// CreatedAt	time.Time `json:"createdAt"`
	CreatedAt 	time.Time `gorm:"column:createdAt"`
}
