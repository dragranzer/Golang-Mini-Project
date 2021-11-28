package response

import "github.com/dragranzer/Golang-Mini-Project/features/users"

type User struct {
	ID       int    `json:"id" form:"id"`
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func FromCore(req users.Core) User {
	return User{
		ID:       req.ID,
		Nama:     req.Nama,
		Email:    req.Email,
		Password: req.Password,
	}
}

func FromCoreSlice(core []users.Core) []User {
	var artArray []User
	for key := range core {
		artArray = append(artArray, FromCore(core[key]))
	}
	return artArray
}
