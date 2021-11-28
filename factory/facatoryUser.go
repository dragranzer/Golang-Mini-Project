package factory

import (
	"github.com/dragranzer/Golang-Mini-Project/config"

	_user_bussiness "github.com/dragranzer/Golang-Mini-Project/features/users/bussiness"
	_user_data "github.com/dragranzer/Golang-Mini-Project/features/users/data"
	_user_presentation "github.com/dragranzer/Golang-Mini-Project/features/users/presentation"
)

type PresenterUser struct {
	UserPresentation _user_presentation.UsersHandler
}

func InitUser() PresenterUser {

	userData := _user_data.NewUserRepository(config.DB)
	userBussiness := _user_bussiness.NewUserBussiness(userData)

	return PresenterUser{
		UserPresentation: *_user_presentation.NewUserHandler(userBussiness),
	}
}
