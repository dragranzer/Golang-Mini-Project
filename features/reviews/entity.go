package reviews

type Core struct {
	UserID     int
	BookID     int
	Review     string
	ListReview []UserRev
}

type UserRev struct {
	UserID int
	Review string
}

type Bussiness interface {
	CreateData(data Core) (err error)
	GetDatabyBookID(id int) (resp Core, err error)
	// continue the function abtraction
}

type Data interface {
	InsertData(data Core) (err error)
	SelectDatabyBookID(id int) (resp []Core, err error)
}
