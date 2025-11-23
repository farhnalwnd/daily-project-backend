package domain

import "gorm.io/gorm"

type ActivityMaster struct {
	gorm.Model

	Activitiy       string `gorm:"type:varchar(100);not null" json:"activity"`
	Description     string `gorm:"type:text" json:"description"`
	IsIncomeSource  bool   `gorm:"not null" json:"is_income_source"`
	IsOutcomeSource bool   `gorm:"not null" json:"is_outcome_source"`

	Activities []Activity `gorm:"foreignKey:ActivityMasterID"`
}
