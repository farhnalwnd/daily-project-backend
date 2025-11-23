package domain

import "gorm.io/gorm"

type ResourceCategory struct {
	gorm.Model

	Category string `gorm:"type:varchar(15);not null" json:"category"`

	Activities []Activity `gorm:"foreignKey:ResourceCategoryID"`
	Budgets    []Budgets  `gorm:"foreignKey:ResourceCategoryID"`
}
