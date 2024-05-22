package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_routes "mkp-cinema-api/app/routes"

	_middlewares "mkp-cinema-api/app/middlewares"
	_userUsecase "mkp-cinema-api/businesses/users"
	_userController "mkp-cinema-api/controllers/users"

	_movieUsecase "mkp-cinema-api/businesses/movies"
	_movieController "mkp-cinema-api/controllers/movies"

	_showtimeUsecase "mkp-cinema-api/businesses/showtimes"
	_showtimeController "mkp-cinema-api/controllers/showtimes"

	_driverFactory "mkp-cinema-api/drivers"
	_dbDriver "mkp-cinema-api/drivers/postgresql"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const DEFAULT_PORT = "3000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}

	db := configDB.InitDB()
	_dbDriver.DBMigrate(db)

	redisSecretKey := os.Getenv("REDIS_SECRET_KEY")
    redisExpiredDurationStr := os.Getenv("REDIS_EXPIRED_DURATION")
    redisAddress := os.Getenv("REDIS_ADDRESS")

    redisExpiredDuration, err := strconv.Atoi(redisExpiredDurationStr)
    if err != nil {
        log.Fatalf("Invalid REDIS_EXPIRED_DURATION: %v", err)
    }

    configJWT := _middlewares.NewConfigJWT(
        redisSecretKey, 
        redisExpiredDuration, 
        redisAddress,
    )

	configLogger := _middlewares.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	app := echo.New()

	app.Use(middleware.Recover())

	userRepo := _driverFactory.NewUserRepository(db)
	userUseCase := _userUsecase.NewUserUsecase(userRepo, configJWT)
	userCtrl := _userController.NewAuthController(userUseCase)

	movieRepo := _driverFactory.NewMovieRepository(db)
	movieUseCase := _movieUsecase.NewMovieUsecase(movieRepo)
	movieCtrl := _movieController.NewMovieController(movieUseCase, configJWT)

	showtimeRepo := _driverFactory.NewShowtimeRepository(db)
	showtimeUseCase := _showtimeUsecase.NewShowtimeUsecase(showtimeRepo)
	showtimeCtrl := _showtimeController.NewShowtimeController(showtimeUseCase, configJWT)

	routesInit := _routes.ControllerList{
		LoggerMiddleware: configLogger.Init(),
		JWTMiddleware:    configJWT.Init(),
		AuthController:   *userCtrl,
		MovieController:  *movieCtrl,
		ShowtimeController: *showtimeCtrl,
	}

	routesInit.RouteRegister(app)

	var port string = os.Getenv("APP_PORT")

	if port == "" {
		port = DEFAULT_PORT
	}

	var appPort string = fmt.Sprintf(":%s", port)

	app.Logger.Fatal(app.Start(appPort))
}
