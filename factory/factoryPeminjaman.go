package factory

// import (
// 	"github.com/dragranzer/Golang-Mini-Project/config"

// 	_peminjaman_bussiness "github.com/dragranzer/Golang-Mini-Project/features/peminjamans/bussiness"
// 	_peminjaman_data "github.com/dragranzer/Golang-Mini-Project/features/peminjamans/data"
// 	_peminjaman_presentation "github.com/dragranzer/Golang-Mini-Project/features/peminjamans/presentation"
// )

// type PresenterPeminjaman struct {
// 	PeminjamanPresentation _peminjaman_presentation.PeminjamansHandler
// }

// func InitPeminjaman() PresenterPeminjaman {

// 	peminjamanData := _peminjaman_data.NewPeminjamanRepository(config.DB)
// 	peminjamanBussiness := _peminjaman_bussiness.NewPeminjamanBussiness(peminjamanData)

// 	return PresenterPeminjaman{
// 		PeminjamanPresentation: *_peminjaman_presentation.NewPeminjamanHandler(peminjamanBussiness),
// 	}
// }
