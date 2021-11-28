package data

import (
	"time"

	"github.com/dragranzer/Golang-Mini-Project/features/users"
)

type User struct {
	ID        int       `json:"id" form:"id"`
	Nama      string    `json:"nama" form:"nama"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

//DTO

func (a *User) toCore() users.Core {
	return users.Core{
		ID:        int(a.ID),
		Nama:      a.Nama,
		Email:     a.Email,
		Password:  a.Password,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func toCoreList(resp []User) []users.Core {
	a := []users.Core{}
	for key := range resp {
		a = append(a, resp[key].toCore())
	}
	return a
}

func fromCore(core users.Core) User {

	return User{
		ID:        int(core.ID),
		Nama:      core.Nama,
		Email:     core.Email,
		Password:  core.Password,
		UpdatedAt: core.UpdatedAt,
	}
}
