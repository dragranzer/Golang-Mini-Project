package bussiness

import "github.com/dragranzer/Golang-Mini-Project/features/users"

type usersUsecase struct {
	userData users.Data
}

func NewUserBussiness(artData users.Data) users.Bussiness {
	return &usersUsecase{
		userData: artData,
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
