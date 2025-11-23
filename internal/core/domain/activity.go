package domain

import "gorm.io/gorm"

type Activity struct {
	gorm.Model

	Type     string `gorm:"type:varchar(50);not null" json:"type"`
	Nominal  int16  `gorm:"not null" json:"nominal"`
	Vendor   string `gorm:"type:varchar(100);not null" json:"vendor"`
	Priority string `gorm:"type:varchar(50);not null" json:"priority"`
	Notes    string `gorm:"type:text" json:"notes"`

	// Foreign Key to ActivityMaster
	ActivityMasterID uint            `gorm:"not null" json:"activity_master_id"`
	Master           *ActivityMaster `gorm:"foreignKey:ActivityMasterID"`

	// Foreign Key to User
	UserID uint  `gorm:"not null" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID"`

	// Foreign Key to ResourceCategory
	ResourceCategoryID uint              `gorm:"not null" json:"resource_category_id"`
	Category           *ResourceCategory `gorm:"foreignKey:ResourceCategoryID"`

	// Foreign Key to Account
	AccountID uint     `gorm:"not null" json:"account_id"`
	Account   *Account `gorm:"foreignKey:AccountID"`
}
