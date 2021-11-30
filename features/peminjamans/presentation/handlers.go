package presentation

import (
	"fmt"
	"net/http"

	"github.com/dragranzer/Golang-Mini-Project/features/peminjamans"
	"github.com/dragranzer/Golang-Mini-Project/features/peminjamans/presentation/request"
	"github.com/dragranzer/Golang-Mini-Project/features/peminjamans/presentation/response"
	"github.com/labstack/echo/v4"
)

type PeminjamansHandler struct {
	peminjamanBussiness peminjamans.Bussiness
}

func NewPeminjamanHandler(ab peminjamans.Bussiness) *PeminjamansHandler {
	return &PeminjamansHandler{
		peminjamanBussiness: ab,
	}
}

func (boh *PeminjamansHandler) GetAllPeminjaman(c echo.Context) error {
	// fmt.Println("Masuk Handlers F2")
	result := boh.peminjamanBussiness.GetAllData()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    response.FromCoreSlice(result),
	})
}

func (ah *PeminjamansHandler) InsertPeminjaman(c echo.Context) error {
	// fmt.Println("Masuk Handlers F2")
	peminjaman := request.Peminjaman{}
	c.Bind(&peminjaman)
	fmt.Println("peminjaman presentation ========== ", peminjaman)
	data, err := ah.peminjamanBussiness.CreateData(request.ToCore(peminjaman))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success",
		"peminjaman": response.FromCore(data),
	})
}

func (ah *PeminjamansHandler) GetDetailPinjam(c echo.Context) error {
	// fmt.Println("Masuk Handlers F2")
	var judul string
	echo.PathParamsBinder(c).String("judul", &judul)
	result := ah.peminjamanBussiness.GetDetailBookPinjam(judul)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    result,
	})
}
