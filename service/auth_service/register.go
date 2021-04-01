package auth_service

import (
	"github.com/EDDYCJY/go-gin-example/models"
)

func Reg(Username, Password string) (bool, error) {
	return models.InsertUser(Username, Password)
}
