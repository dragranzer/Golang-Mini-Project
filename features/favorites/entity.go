package favorites

type Core struct {
	UserID   int
	BookID   int
	UserName []string
	BookName []string
}

type Bussiness interface {
	CreateData(data Core) (err error)
	GetDatabyUserID(id int) (resp Core, err error)
	GetDatabyBookID(id int) (resp Core, err error)
	// continue the function abtraction
}

type Data interface {
	InsertData(data Core) (err error)
	SelectDatabyUserID(id int) (resp []Core, err error)
	SelectDatabyBookID(id int) (resp []Core, err error)
}
