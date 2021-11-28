package data

import (
	"time"

	"github.com/dragranzer/Golang-Mini-Project/features/users"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id" form:"id"`
	Nama      string    `json:"nama" form:"nama"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

//DTO

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *User) toCore() users.Core {
	return users.Core{
		ID:        int(a.ID),
		Nama:      a.Nama,
		Email:     a.Email,
		Password:  a.Password,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func toCoreList(resp []User) []users.Core {
	a := []users.Core{}
	for key := range resp {
		a = append(a, resp[key].toCore())
	}
	return a
}

func fromCore(core users.Core) User {

	password, _ := HashPassword(core.Password)
	return User{
		Nama:     core.Nama,
		Email:    core.Email,
		Password: password,
	}
}
