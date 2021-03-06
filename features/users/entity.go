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

type UserResp struct {
	ID    int    `json:"id" form:"id"`
	Nama  string `json:"nama" form:"nama"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}

type Bussiness interface {
	CreateData(data Core) (err error)
	GetAllData() (resp []Core)
	GetDatabyName(name string) (resp Core, err error)
	GetDatabyID(id int) (resp Core, err error)
	ChangeDatabyID(id int, newData Core) (err error)
	Login(email string, pass string) (resp UserResp, isAuth bool, err error)
	// continue the function abtraction
}

type Data interface {
	InsertData(data Core) (err error)
	SelectAllData() (resp []Core)
	SelectDatabyName(name string) (resp Core, err error)
	SelectDatabyID(id int) (resp Core, err error)
	SelectDatabyEmail(email string) (resp Core, err error)
	UpdateDatabyID(id int, newData Core) (err error)
	CheckEmailPass(email string, pass string) (isAuth bool, user Core, err error)
}
