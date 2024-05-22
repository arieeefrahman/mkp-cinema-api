package showtimes

import (
	"fmt"
	"mkp-cinema-api/app/middlewares"
	"mkp-cinema-api/businesses/showtimes"
	ctrl "mkp-cinema-api/controllers"
	"mkp-cinema-api/controllers/showtimes/request"
	"mkp-cinema-api/controllers/showtimes/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ShowtimeController struct {
	showtimeUsecase showtimes.Usecase
	jwtConfig    *middlewares.ConfigJWT // Add this field to use ConfigJWT
}

func NewShowtimeController(showtimeUc showtimes.Usecase, jwtConfig *middlewares.ConfigJWT) *ShowtimeController { // Update constructor to accept ConfigJWT
	return &ShowtimeController{
		showtimeUsecase: showtimeUc,
		jwtConfig:    jwtConfig,
	}
}

func (showtimeCtrl *ShowtimeController) GetAll(c echo.Context) error {
	showtimeData, err := showtimeCtrl.showtimeUsecase.GetAll()
	if err != nil {
		return ctrl.NewErrorResponse(c, 
			http.StatusInternalServerError, 
			"failed", "internal server error", 
			fmt.Sprintf("%s", err),
		)
	}

	showtimes := []response.Showtime{}

	for _, showtime := range showtimeData {
		showtimes = append(showtimes, response.FromDomain(showtime))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "all showtime data", showtimes)
}

func (showtimeCtrl *ShowtimeController) GetByID(c echo.Context) error {
	id := c.Param("id")

	showtime, err := showtimeCtrl.showtimeUsecase.GetByID(id)
	if err != nil {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", fmt.Sprintf("%s", err))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "showtime found", response.FromDomain(showtime))
}

func (showtimeCtrl *ShowtimeController) Create(c echo.Context) error {
	input := request.ShowtimeCreate{}
	if err := c.Bind(&input); err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "validation failed", fmt.Sprintf("%s", err))
	}

	err := input.Validate()
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "validation failed", fmt.Sprintf("%s", err))
	}

	showtime, err := showtimeCtrl.showtimeUsecase.Create(input.ToDomain())
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "create failed", fmt.Sprintf("%s", err))
	}

	return ctrl.NewResponse(c, http.StatusCreated, "success", "showtime created", response.FromDomain(showtime))
}


func (showtimeCtrl *ShowtimeController) Update(c echo.Context) error {
	input := request.ShowtimeUpdate{}
	if err := c.Bind(&input); err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "validation failed", fmt.Sprintf("%s", err))
	}

	showtimeId := c.Param("id")

	err := input.Validate()
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "validation failed", fmt.Sprintf("%s", err))
	}

	showtime, err := showtimeCtrl.showtimeUsecase.Update(showtimeId, input.ToDomain())
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusNotFound, "failed", "update failed", fmt.Sprintf("%s", err))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "showtime updated", response.FromDomain(showtime))
}

func (showtimeCtrl *ShowtimeController) Delete(c echo.Context) error {
	showtimeId := c.Param("id")
	_, err := showtimeCtrl.showtimeUsecase.Delete(showtimeId)
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusNotFound, "failed", "delete failed", fmt.Sprintf("%s", err))
	}

	return ctrl.NewInfoResponse(c, http.StatusOK, "success", "showtime deleted")
}
