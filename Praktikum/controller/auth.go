package controller

import (
	"Praktikum/model"
	"Praktikum/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	LoginUserController(c echo.Context) error
}

type authController struct {
	userUsecase usecase.UserUsecase
}

func NewAuthController(userUsecase usecase.UserUsecase) *authController {
	return &authController{
		userUsecase,
	}
}

func (a *authController) LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := a.userUsecase.LoginUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    user,
	})
}
