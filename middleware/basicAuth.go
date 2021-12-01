package middleware

import (
	"github.com/dragranzer/Golang-Mini-Project/config"
	"github.com/dragranzer/Golang-Mini-Project/features/users/data"
	"github.com/labstack/echo/v4"
)

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	var user = data.User{}
	err := config.DB.Where("email = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
