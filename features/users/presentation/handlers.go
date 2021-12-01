package presentation

import (
	"fmt"
	"net/http"

	"github.com/dragranzer/Golang-Mini-Project/features/users"
	"github.com/dragranzer/Golang-Mini-Project/features/users/presentation/request"
	"github.com/dragranzer/Golang-Mini-Project/features/users/presentation/response"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	userBussiness users.Bussiness
}

func NewUserHandler(ab users.Bussiness) *UsersHandler {
	return &UsersHandler{
		userBussiness: ab,
	}
}

func (boh *UsersHandler) GetAllUser(c echo.Context) error {
	result := boh.userBussiness.GetAllData()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "handlers user",
		"data":    response.FromCoreSlice(result),
	})
}

func (ah *UsersHandler) InsertUser(c echo.Context) error {
	user := request.User{}
	c.Bind(&user)
	err := ah.userBussiness.CreateData(request.ToCore(user))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "data success di masukkan",
	})
}

func (uH *UsersHandler) GetUser(c echo.Context) error {
	fmt.Println("masuk handlers")
	var name string
	echo.PathParamsBinder(c).String("nama", &name)
	fmt.Println(name)
	// result := []users.Core{}
	// result = append(result, uH.userBussiness.GetDatabyName(name))
	result, _ := uH.userBussiness.GetDatabyName(name)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    result,
	})
}

func (uH *UsersHandler) UpdateUserData(c echo.Context) error {
	user := request.User{}
	var idstring string
	echo.PathParamsBinder(c).String("id", &idstring)
	c.Bind(&user)
	// id, _ := strconv.Atoi(idstring)
	fmt.Println("id ", user.ID)
	fmt.Println("newdata", user)
	err := uH.userBussiness.ChangeDatabyID(user.ID, request.ToCore(user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data changed",
	})
}

func (uH *UsersHandler) LoginUser(c echo.Context) error {
	user := request.User{}
	c.Bind(&user)
	resp, isAuth, err := uH.userBussiness.Login(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if !isAuth {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Unauthorized",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data changed",
		"data":    resp,
	})
}
