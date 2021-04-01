package models

import (
	"errors"
)

type userinfo struct {
	ID       int    `gorm:"AUTO_INCREMENT"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// 用户注册
func InsertUser(username, password string) (bool, error) {
	var user userinfo
	user = userinfo{Username: username, Password: password}

	if result := db.Where("username = ?", username).First(&user); result.RowsAffected > 0 {
		return false, errors.New("用户名重复")
	}

	if err := db.Create(&user).Error; err != nil {
		return false, err
	}
	return false, nil
}
