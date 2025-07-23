package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;default:null"`
}

// func (user *User) BeforeCreate(tx *gorm.DB) {
// 	if user.Email == "" {
// 		user.Email == nil
// 	}
// 	return
// }
