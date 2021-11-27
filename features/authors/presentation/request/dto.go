package request

import "github.com/dragranzer/Golang-Mini-Project/features/authors"

type Author struct {
	Nama string `json:"nama" form:"nama"`
}

func ToCore(req Author) authors.Core {
	return authors.Core{
		Nama: req.Nama,
	}
}
