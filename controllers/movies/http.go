package movies

import (
	"fmt"
	"mkp-cinema-api/app/middlewares"
	"mkp-cinema-api/businesses/movies"
	ctrl "mkp-cinema-api/controllers"
	"mkp-cinema-api/controllers/movies/request"
	"mkp-cinema-api/controllers/movies/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MovieController struct {
	movieUsecase movies.Usecase
	jwtConfig    *middlewares.ConfigJWT // Add this field to use ConfigJWT
}

func NewMovieController(movieUc movies.Usecase, jwtConfig *middlewares.ConfigJWT) *MovieController { // Update constructor to accept ConfigJWT
	return &MovieController{
		movieUsecase: movieUc,
		jwtConfig:    jwtConfig,
	}
}

func (mc *MovieController) GetAll(c echo.Context) error {
	moviesData, err := mc.movieUsecase.GetAll()
	if err != nil {
		return ctrl.NewErrorResponse(c, 
			http.StatusInternalServerError, 
			"failed", "internal server error", 
			fmt.Sprintf("%s", err),
		)
	}

	movies := []response.Movie{}

	for _, movie := range moviesData {
		movies = append(movies, response.FromDomain(movie))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "all movies", movies)
}

func (mc *MovieController) GetByID(c echo.Context) error {
	id := c.Param("id")

	movie, err := mc.movieUsecase.GetByID(id)
	if err != nil {
		return ctrl.NewInfoResponse(c, http.StatusBadRequest, "failed", fmt.Sprintf("%s", err))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "movie found", response.FromDomain(movie))
}

func (mc *MovieController) Create(c echo.Context) error {
	input := request.MovieCreate{}
	if err := c.Bind(&input); err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "validation failed", fmt.Sprintf("%s", err))
	}

	err := input.Validate()
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "validation failed", fmt.Sprintf("%s", err))
	}

	movie, err := mc.movieUsecase.Create(input.ToDomain())
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "create failed", fmt.Sprintf("%s", err))
	}

	return ctrl.NewResponse(c, http.StatusCreated, "success", "movie created", response.FromDomain(movie))
}


func (mc *MovieController) Update(c echo.Context) error {
	input := request.MovieUpdate{}
	if err := c.Bind(&input); err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "validation failed", fmt.Sprintf("%s", err))
	}

	movieId := c.Param("id")

	err := input.Validate()
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusBadRequest, "failed", "validation failed", fmt.Sprintf("%s", err))
	}

	movie, err := mc.movieUsecase.Update(movieId, input.ToDomain())
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusNotFound, "failed", "update failed", fmt.Sprintf("%s", err))
	}

	return ctrl.NewResponse(c, http.StatusOK, "success", "movie updated", response.FromDomain(movie))
}

func (mc *MovieController) Delete(c echo.Context) error {
	movieId := c.Param("id")
	_, err := mc.movieUsecase.Delete(movieId)
	if err != nil {
		return ctrl.NewErrorResponse(c, http.StatusInternalServerError, "failed", "delete failed", fmt.Sprintf("%s", err))
	}

	return ctrl.NewInfoResponse(c, http.StatusOK, "success", "movie deleted")
}
