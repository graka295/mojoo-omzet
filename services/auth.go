package services

import (
	"mojoo/omzet/container"
	"mojoo/omzet/model/domain"
	"mojoo/omzet/model/dto"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// AuthServices contract
type AuthServices interface {
	Login(c echo.Context) error
}

// AuthServicesImpl struct authServices
type AuthServicesImpl struct {
	container container.Container
}

func NewAuthServices(container container.Container) AuthServices {
	return AuthServicesImpl{
		container: container,
	}
}
func (x AuthServicesImpl) Login(c echo.Context) error {
	req := dto.LoginRequest{}
	if err := c.Bind(&req); err != nil {
		return x.container.GetResponse().InternalServerError(c, err.Error())
	}
	messageValidate, valid := x.container.GetValidate().Validate(req)
	if !valid {
		return x.container.GetResponse().BadRequest(c, messageValidate)
	}

	users := domain.NewUsers()
	data := users.Login(x.container, req.Username, req.Password)
	if data != nil {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = data.Name
		claims["id"] = data.ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		dto := dto.LoginResponse{
			Name:     data.Name,
			Username: data.Username,
			Token:    t,
		}
		return x.container.GetResponse().Ok(c, dto)
	}
	return x.container.GetResponse().Unauthorized(c, map[string]string{
		"address": "account",
		"message": "Username or password is wrong",
	})
}
