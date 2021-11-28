package users

import "time"

type Core struct {
	ID        int       `json:"id" form:"id"`
	Nama      string    `json:"nama" form:"nama"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type Bussiness interface {
	CreateData(data Core) (err error)
	GetAllData() (resp []Core)
	// continue the function abtraction
}

type Data interface {
	InsertData(data Core) (err error)
	SelectAllData() (resp []Core)
}