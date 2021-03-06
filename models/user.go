package models

import (
	"log"
)

type User struct {
	Model

	Name     string `gorm:"size:100;comment:用户名" json:"name"`
	Password string `gorm:"size:255;comment:密码" json:"password"`
	Avatar   string `gorm:"size:512;comment:头像" json:"avatar"`
	Email    string `gorm:"size:255;comment:邮箱" json:"email"`
	Phone    string `gorm:"size:50;comment:手机号" json:"phone"`
	Type     int    `gorm:"comment:用户角色 -1 管理员 1 普通用户" json:"type"`
	Status   int    `gorm:"comment:状态 1 正常 0 待激活 -1 禁用" json:"status"`
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateAdmin() bool {
	var admin User
	adminEmail := "wwc@admin.com"
	result := db.Where("email = ?", adminEmail).First(&admin)
	if result.Error != nil {
		admin.Name = "系统管理员"
		admin.Email = "wwc@admin.com"
		admin.Password = "123456"
		admin.Type = -1
		admin.Status = 1

		result := db.Create(&admin)
		if result.Error != nil {
			log.Fatalf("create admin err: %v", result.Error)
			return false
		}
	}
	return true
}
