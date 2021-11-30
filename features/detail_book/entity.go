package detail_book

type Core struct {
	BookID   int `gorm:"primaryKey"`
	AuthorID int `gorm:"primaryKey"`
}

type Bussiness interface {
	CreateData(data Core) (resp Core, err error)
	GetAuthorbyBookID(id int) (authorID []int)
}

type Data interface {
	InsertData(data Core) (resp Core, err error)
	SelectAuthorbyBookID(id int) (authorID []int)
}
