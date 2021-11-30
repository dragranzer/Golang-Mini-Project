package bussiness

import (
	"github.com/dragranzer/Golang-Mini-Project/features/authors"
)

type authorsUsecase struct {
	authorData authors.Data
}

func NewAuthorBussiness(auData authors.Data) authors.Bussiness {
	return &authorsUsecase{
		authorData: auData,
	}
}

func (bu *authorsUsecase) CreateData(data authors.Core) (resp authors.Core, err error) {

	data, err = bu.authorData.InsertData(data)
	if err != nil {
		return authors.Core{}, err
	}
	return data, nil
}

func (bu *authorsUsecase) GetAllData() (resp []authors.Core) {
	resp = bu.authorData.SelectAllData()
	return
}

func (bu *authorsUsecase) GetDetailData(judul string) (resp []authors.Core, err error) {
	resp, err = bu.authorData.SelectData(judul)
	// fmt.Println("get detail data")
	if err != nil {
		return []authors.Core{}, err
	}

	return resp, nil
}

func (bu *authorsUsecase) GetDetailDatabyID(id int) (resp authors.Core, err error) {
	resp, err = bu.authorData.SelectDatabyID(id)
	// fmt.Println("get detail data")
	if err != nil {
		return authors.Core{}, err
	}

	return resp, nil
}
