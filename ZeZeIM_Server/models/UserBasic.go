package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name       string `gorm:"type:varchar(20);not null"`
	Password   string `gorm:"type:varchar(20);not null"`
	Phone      string `gorm:"type:varchar(20);not null"`
	Sexy       string `gorm:"type:varchar(20);not null"`
	Email      string `gorm:"type:varchar(20);not null"`
	ClientIP   string `gorm:"type:varchar(20);not null"`
	Identity   string `gorm:"type:varchar(20);not null"`
	ClientPort string `gorm:"type:varchar(20);not null"`
	LoginTime  uint64 `gorm:"type:bigint(20);not null"`
	LogoutTime uint64 `gorm:"type:bigint(20);not null"`
	IsLogout   bool   `gorm:"type:boolean;not null"`
	DeviceInfo string `gorm:"type:varchar(20);not null"`
}

func (UserBasic) TableName() string {
	return "user_basic"
}
