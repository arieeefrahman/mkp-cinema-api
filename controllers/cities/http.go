package cities

import (
	"fmt"
	"mkp-cinema-api/app/middlewares"
	"mkp-cinema-api/businesses/cities"
	ctrl "mkp-cinema-api/controllers"
	"mkp-cinema-api/controllers/cities/request"
	"mkp-cinema-api/controllers/cities/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CityController struct {
	cityUsecase cities.Usecase
	jwtConfig    *middlewares.ConfigJWT
}

func NewCityController(cityUc cities.Usecase, jwtConfig *middlewares.ConfigJWT) *CityController { 
	return &CityController{
		cityUsecase: cityUc,
		jwtConfig:    jwtConfig,
	}
}

func (cityCtrl *CityController) GetAll(c echo.Context) error {
	citiesData, err := cityCtrl.cityUsecase.GetAll()
	if err != nil {
		return ctrl.NewErrorResponse(c, 
			http.StatusInternalServerError, 
			"failed", "internal server error", 
			fmt.Sprintf("%s", err),
		)
	}

	cities := []response.City{}

	for _, city := range citiesData {
		cities = append(cities, response.FromDomain(city))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "all cities", cities)
}

func (cityCtrl *CityController) GetByID(c echo.Context) error {
	id := c.Param("id")

	city, err := cityCtrl.cityUsecase.GetByID(id)
	if err != nil {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", fmt.Sprintf("%s", err))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "city found", response.FromDomain(city))
}

func (cityCtrl *CityController) Create(c echo.Context) error {
	input := request.CityCreate{}
	if err := c.Bind(&input); err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "validation failed", fmt.Sprintf("%s", err))
	}

	err := input.Validate()
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "validation failed", fmt.Sprintf("%s", err))
	}

	city, err := cityCtrl.cityUsecase.Create(input.ToDomain())
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "create failed", fmt.Sprintf("%s", err))
	}

	return ctrl.NewResponse(c, http.StatusCreated, "success", "city created", response.FromDomain(city))
}


func (cityCtrl *CityController) Update(c echo.Context) error {
	input := request.CityUpdate{}
	if err := c.Bind(&input); err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "validation failed", fmt.Sprintf("%s", err))
	}

	cityId := c.Param("id")

	err := input.Validate()
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "validation failed", fmt.Sprintf("%s", err))
	}

	city, err := cityCtrl.cityUsecase.Update(cityId, input.ToDomain())
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusNotFound, "failed", "update failed", fmt.Sprintf("%s", err))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "city updated", response.FromDomain(city))
}

func (cityCtrl *CityController) Delete(c echo.Context) error {
	cityId := c.Param("id")
	_, err := cityCtrl.cityUsecase.Delete(cityId)
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusInternalServerError, "failed", "delete failed", fmt.Sprintf("%s", err))
	}

	return ctrl.NewInfoResponse(c, http.StatusOK, "success", "city deleted")
}
