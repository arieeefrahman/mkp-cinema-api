package users

import (
	"fmt"
	"net/http"

	"mkp-cinema-api/businesses/users"

	ctrl "mkp-cinema-api/controllers"
	"mkp-cinema-api/controllers/users/request"
	"mkp-cinema-api/controllers/users/response"

	"github.com/labstack/echo/v4"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

type AuthController struct {
	authUsecase users.Usecase
}

func NewAuthController(authUC users.Usecase) *AuthController {
	return &AuthController{
		authUsecase: authUC,
	}
}

func (ac *AuthController) HelloMessage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello there!")
}

func (ac *AuthController) Register(c echo.Context) error {
	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", fmt.Sprintf("%s", err))
	}

	err := userInput.Validate()

	if err != nil {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", fmt.Sprintf("%s", err))
	}

	if userInput.Password != userInput.ConfirmationPassword {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", "password and confirmation password do not match")
	}

	const minEntropyBits = 30
	err = passwordvalidator.Validate(userInput.Password, minEntropyBits)
	
	if err != nil {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", fmt.Sprintf("%s", err))
	}

	user, err := ac.authUsecase.Register(userInput.ToDomainRegister())

	if err != nil {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", fmt.Sprintf("%s", err))
	}

	return ctrl.NewResponse(c, http.StatusCreated, "success", "account created", response.FromDomain(user))
}

func (ac *AuthController) Login(c echo.Context) error {
	userInput := request.UserLogin{}

	if err := c.Bind(&userInput); err != nil {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", fmt.Sprintf("%s", err))
	}

	err := userInput.Validate()
	if err != nil {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", fmt.Sprintf("%s", err))
	}

	token, err := ac.authUsecase.Login(userInput.ToDomainLogin())
	if err != nil {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", "invalid email or password")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"access_token": token["access_token"],
	})
}