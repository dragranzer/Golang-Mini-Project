package response

import "github.com/dragranzer/Golang-Mini-Project/features/authors"

type Author struct {
	ID   int    `json:"id" form:"id"`
	Nama string `json:"nama" form:"nama"`
}

func FromCore(res authors.Core) Author {
	return Author{
		ID:   res.ID,
		Nama: res.Nama,
	}
}

func FromCoreSlice(core []authors.Core) []Author {
	var artArray []Author
	for key := range core {
		artArray = append(artArray, FromCore(core[key]))
	}
	return artArray
}
