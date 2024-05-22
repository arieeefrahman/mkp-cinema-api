package drivers

import (
	userDomain "mkp-cinema-api/businesses/users"
	userDB "mkp-cinema-api/drivers/postgresql/users"

	movieDomain "mkp-cinema-api/businesses/movies"
	movieDB "mkp-cinema-api/drivers/postgresql/movies"

	cityDomain "mkp-cinema-api/businesses/cities"
	cityDB "mkp-cinema-api/drivers/postgresql/cities"

	showtimeDomain "mkp-cinema-api/businesses/showtimes"
	showtimeDB "mkp-cinema-api/drivers/postgresql/showtimes"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}

func NewMovieRepository(conn *gorm.DB) movieDomain.Repository {
	return movieDB.NewMySQLRepository(conn)
}

func NewCityRepository(conn *gorm.DB) cityDomain.Repository {
	return cityDB.NewMySQLRepository(conn)
}

func NewShowtimeRepository(conn *gorm.DB) showtimeDomain.Repository {
	return showtimeDB.NewMySQLRepository(conn)
}
