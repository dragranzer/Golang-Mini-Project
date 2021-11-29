package bussiness

import (
	"github.com/dragranzer/Golang-Mini-Project/features/peminjamans"
)

type peminjamansUsecase struct {
	peminjamanData peminjamans.Data
}

func NewPeminjamanBussiness(artData peminjamans.Data) peminjamans.Bussiness {
	return &peminjamansUsecase{
		peminjamanData: artData,
	}
}

func (bu *peminjamansUsecase) CreateData(data peminjamans.Core) (resp peminjamans.Core, err error) {
	// if err := bu.validate.Struct(data); err != nil {
	// 	return peminjamans.Core{}, err
	// }

	data, err = bu.peminjamanData.InsertData(data)
	if err != nil {
		return peminjamans.Core{}, err
	}
	return data, nil
}

func (bu *peminjamansUsecase) GetAllData() (resp []peminjamans.Core) {
	resp = bu.peminjamanData.SelectAllData()
	return resp
}
