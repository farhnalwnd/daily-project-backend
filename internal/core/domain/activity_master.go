package domain

import "gorm.io/gorm"

type ActivityMaster struct {
	gorm.Model

	Activitiy   string `gorm:"type:varchar(100);not null" json:"activity"`
	Description string `gorm:"type:text" json:"description"`

	Activities []Activity `gorm:"foreignKey:ActivityMasterID"`
}
