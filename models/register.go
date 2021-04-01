package models

type Reg struct {
	ID       int    `gorm:"AUTO_INCREMENT"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckReg Confirm registration information
func insertUser(username, password string) (bool, error) {
	var user Reg
	user = Reg{Username: username, Password: password}
	result := db.Create(&user)

	print(result)
	//err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	return false, err
	//}
	//
	//if auth.ID > 0 {
	//	return true, nil
	//}

	return false, nil
}
