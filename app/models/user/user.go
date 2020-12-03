package user

import (
	"goblog/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`
	// gorm:"_" —— 设置 GORM 在读写时略过此字段
	PasswordConfirm string `  gorm:"_" valid:"password_confirm"`
}
