package routes

import (
	"mkp-cinema-api/controllers/cities"
	"mkp-cinema-api/controllers/movies"
	"mkp-cinema-api/controllers/showtimes"
	"mkp-cinema-api/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      middleware.JWTConfig
	AuthController     users.AuthController
	MovieController    movies.MovieController
	CityController     cities.CityController
	ShowtimeController showtimes.ShowtimeController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderServer,
		},
	}))

	e.Use(cl.LoggerMiddleware)

	v1 := e.Group("/api/v1")
	v1.GET("", cl.AuthController.HelloMessage)
	v1.POST("/register", cl.AuthController.Register)
	v1.POST("/login", cl.AuthController.Login)

	movies := v1.Group("/movies")
	movies.GET("", cl.MovieController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))
	movies.GET("/:id", cl.MovieController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))
	movies.POST("", cl.MovieController.Create, middleware.JWTWithConfig(cl.JWTMiddleware))
	movies.PUT("/:id", cl.MovieController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	movies.DELETE("/:id", cl.MovieController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))

	cities := v1.Group("/cities")
	cities.GET("", cl.CityController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))
	cities.GET("/:id", cl.CityController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))
	cities.POST("", cl.CityController.Create, middleware.JWTWithConfig(cl.JWTMiddleware))
	cities.PUT("/:id", cl.CityController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	cities.DELETE("/:id", cl.CityController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))

	showtime := v1.Group("/showtime")
	showtime.GET("", cl.ShowtimeController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))
	showtime.GET("/:id", cl.ShowtimeController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))
	showtime.POST("", cl.ShowtimeController.Create, middleware.JWTWithConfig(cl.JWTMiddleware))
	showtime.PUT("/:id", cl.ShowtimeController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	showtime.DELETE("/:id", cl.ShowtimeController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))
}
