package bussiness

import (
	"github.com/dragranzer/Golang-Mini-Project/features/books"
	"github.com/dragranzer/Golang-Mini-Project/features/users"
)

type usersUsecase struct {
	userData      users.Data
	bookBussiness books.Bussiness
}

func NewUserBussiness(usrData users.Data, bB books.Bussiness) users.Bussiness {
	return &usersUsecase{
		userData:      usrData,
		bookBussiness: bB,
	}
}

func (bu *usersUsecase) CreateData(data users.Core) (err error) {
	// if err := bu.validate.Struct(data); err != nil {
	// 	return users.Core{}, err
	// }

	err = bu.userData.InsertData(data)
	if err != nil {
		return err
	}
	return nil
}

func (bu *usersUsecase) GetAllData() (resp []users.Core) {
	resp = bu.userData.SelectAllData()
	return resp
}

func (bu *usersUsecase) GetDatabyName(name string) (resp users.Core, err error) {
	resp, _ = bu.userData.SelectDatabyName(name)
	return resp, nil
}
