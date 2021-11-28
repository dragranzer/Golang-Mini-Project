package request

import "github.com/dragranzer/Golang-Mini-Project/features/users"

type User struct {
	ID       int    `json:"id" form:"id"`
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(req User) users.Core {
	return users.Core{
		ID:       req.ID,
		Nama:     req.Nama,
		Email:    req.Email,
		Password: req.Password,
	}
}
