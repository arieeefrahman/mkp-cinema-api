package postgresql_driver

import (
	"fmt"
	"log"
	"mkp-cinema-api/drivers/postgresql/cities"
	"mkp-cinema-api/drivers/postgresql/movies"
	"mkp-cinema-api/drivers/postgresql/showtimes"
	"mkp-cinema-api/drivers/postgresql/users"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func (config *ConfigDB) InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to database: %s", err)
	}

	log.Println("connected to database")

	return db
}

func DBMigrate(db *gorm.DB) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	db.AutoMigrate(
		&users.User{},
		&movies.Movie{}, 
		&cities.City{},
		&showtimes.Showtime{},
	)
}
