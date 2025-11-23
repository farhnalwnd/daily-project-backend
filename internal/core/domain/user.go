package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null; hash" json:"-"`
	Age      uint8  `gorm:"not null" json:"age"`

	// relations
	Activities []Activity `gorm:"foreignKey:UserID"`
	Accounts   []Account  `gorm:"foreignKey:UserID"`
}
