package domain

import (
	"time"

	"gorm.io/gorm"
)

type Budgets struct {
	gorm.Model

	// fk to User
	UserID uint  `gorm:"not null" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID"`

	// fk to ResourceCategory
	ResourceCategoryID uint              `gorm:"not null" json:"resource_category_id"`
	Category           *ResourceCategory `gorm:"foreignKey:ResourceCategoryID"`

	Amount     int64     `gorm:"not null" json:"amount"`
	StartDate  time.Time `gorm:"not null" json:"start_date"`
	EndDate    time.Time `gorm:"not null" json:"end_date"`
	PeriodType string    `gorm:"type:varchar(50);not null" json:"period_type"`
}
