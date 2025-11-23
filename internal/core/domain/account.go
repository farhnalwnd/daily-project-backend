package domain

import "gorm.io/gorm"

type Account struct {
	gorm.Model

	Name           string `gorm:"type:varchar(100);not null" json:"name"`
	InitialBalance int64  `gorm:"not null" json:"initial_balance"`
	CurrentBalance int64  `gorm:"not null" json:"current_balance"`

	// Foreign Key to User
	UserID uint  `gorm:"not null" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID"`
}
