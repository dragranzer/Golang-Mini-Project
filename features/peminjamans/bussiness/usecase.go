package bussiness

import (
	"fmt"

	"github.com/dragranzer/Golang-Mini-Project/features/books"
	"github.com/dragranzer/Golang-Mini-Project/features/peminjamans"
	"github.com/dragranzer/Golang-Mini-Project/features/users"
)

type peminjamansUsecase struct {
	peminjamanData peminjamans.Data
	bookBussiness  books.Bussiness
	userBussiness  users.Bussiness
}

func NewPeminjamanBussiness(artData peminjamans.Data, bB books.Bussiness, uB users.Bussiness) peminjamans.Bussiness {
	return &peminjamansUsecase{
		peminjamanData: artData,
		bookBussiness:  bB,
		userBussiness:  uB,
	}
}

func (bu *peminjamansUsecase) CreateData(data peminjamans.Core) (resp peminjamans.Core, err error) {
	// if err := bu.validate.Struct(data); err != nil {
	// 	return peminjamans.Core{}, err
	// }
	book, _ := bu.bookBussiness.GetDetailData(data.BookName)
	userid, _ := bu.userBussiness.GetDatabyName(data.UserName)
	insertdata := peminjamans.Core{
		ID:         data.ID,
		Hari:       data.Hari,
		TotalHarga: data.Hari * book[0].Harga,
		BookID:     book[0].ID,
		UserID:     userid.ID,
	}
	data, err = bu.peminjamanData.InsertData(insertdata)
	if err != nil {
		return peminjamans.Core{}, err
	}
	// fmt.Println(book[0].Judul)
	_, _ = bu.bookBussiness.ChangeTersediabyName(book[0].Judul)
	return data, nil
}

func (bu *peminjamansUsecase) GetAllData() (resp []peminjamans.Core) {
	resp = bu.peminjamanData.SelectAllData()
	return resp
}

func (bu *peminjamansUsecase) GetDetailBookPinjam(judul string) (resp peminjamans.DetailBookPinjam) {
	book, _ := bu.bookBussiness.GetDetailData(judul)
	fmt.Println(book)
	listPeminjaman := bu.peminjamanData.SelectDetailBookPinjam(book[0].ID)
	fmt.Println("list peminjaman ", listPeminjaman)
	listPeminjamID := []peminjamans.Peminjam{}
	for _, value := range listPeminjaman {
		temp, _ := bu.userBussiness.GetDatabyID(value.UserID)
		listPeminjamID = append(listPeminjamID, peminjamans.Peminjam{
			ID:   value.UserID,
			Nama: temp.Nama,
		})
	}
	resp = peminjamans.DetailBookPinjam{
		ID:        book[0].ID,
		Peminjams: listPeminjamID,
	}

	return resp
}
