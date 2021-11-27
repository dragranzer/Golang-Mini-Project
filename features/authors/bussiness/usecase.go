package bussiness

import "github.com/dragranzer/Golang-Mini-Project/features/authors"

type authorsUsecase struct {
	authorData authors.Data
}

func NewAuthorBussiness(artData authors.Data) authors.Bussiness {
	return &authorsUsecase{
		authorData: artData,
	}
}

func (bu *authorsUsecase) CreateData(data authors.Core) (resp authors.Core, err error) {
	// if err := bu.validate.Struct(data); err != nil {
	// 	return authors.Core{}, err
	// }

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
