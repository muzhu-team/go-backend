package models

type userinfo struct {
	ID       int    `gorm:"AUTO_INCREMENT"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// 用户注册
func InsertUser(username, password string) (bool, error) {
	var user userinfo
	user = userinfo{Username: username, Password: password}
	if err := db.Create(&user).Error; err != nil {
		return false, err
	}
	//TODO 需要查询一下是否有这个用户，如果有返回错误

	return false, nil
}
