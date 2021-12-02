package bussiness

import (
	"time"

	"github.com/dragranzer/Golang-Mini-Project/constant"
	"github.com/dragranzer/Golang-Mini-Project/features/books"
	"github.com/dragranzer/Golang-Mini-Project/features/users"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type usersUsecase struct {
	userData      users.Data
	bookBussiness books.Bussiness
}

func NewUserBussiness(usrData users.Data, bB books.Bussiness) users.Bussiness {
	return &usersUsecase{
		userData:      usrData,
		bookBussiness: bB,
	}
}

func (bu *usersUsecase) CreateData(data users.Core) (err error) {
	// if err := bu.validate.Struct(data); err != nil {
	// 	return users.Core{}, err
	// }

	err = bu.userData.InsertData(data)
	if err != nil {
		return err
	}
	return nil
}

func (bu *usersUsecase) GetAllData() (resp []users.Core) {
	resp = bu.userData.SelectAllData()
	return resp
}

func (bu *usersUsecase) GetDatabyName(name string) (resp users.Core, err error) {
	resp, _ = bu.userData.SelectDatabyName(name)
	return resp, nil
}

func (bu *usersUsecase) GetDatabyID(id int) (resp users.Core, err error) {
	resp, _ = bu.userData.SelectDatabyID(id)
	return resp, nil
}

func (bu *usersUsecase) ChangeDatabyID(id int, newData users.Core) (err error) {
	err = bu.userData.UpdateDatabyID(id, newData)

	return err
}

func (bu *usersUsecase) Login(email string, pass string) (resp users.UserResp, isAuth bool, err error) {
	// hashpass, _ := migrate.HashPassword(pass)
	// isAuth, user, _ := bu.userData.CheckEmailPass(email, hashpass)

	cekUser, _ := bu.userData.SelectDatabyEmail(email)
	err = bcrypt.CompareHashAndPassword([]byte(cekUser.Password), []byte(pass))
	encryptionErr := (err == nil)

	// if !encryptionErr {
	// 	return resp, false, nil
	// }

	// token, err := middleware.CreateToken(cekUser.ID, cekUser.Nama)

	claims := jwt.MapClaims{
		"userid": cekUser.ID,
		"name":   cekUser.Nama,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	}
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenWithClaims.SignedString([]byte(constant.SECRET_JWT))
	// fmt.Println("token2", token)
	// fmt.Println("ispass? ", encryptionErr)

	if err != nil {
		return resp, false, err
	}

	resp = users.UserResp{
		ID:    cekUser.ID,
		Nama:  cekUser.Nama,
		Email: cekUser.Email,
		Token: token,
	}
	// fmt.Println(resp)
	return resp, encryptionErr, err
}
