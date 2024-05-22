package cinemas

import (
	"errors"
	"mkp-cinema-api/businesses/cinemas"

	"gorm.io/gorm"
)

type cinemaRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) cinemas.Repository {
	return &cinemaRepository{
		conn: conn,
	}
}

func (cinemaRepo *cinemaRepository) GetAll() ([]cinemas.Domain, error) {
	var rec []Cinema

	if err := cinemaRepo.conn.Find(&rec).Error; err != nil {
		return nil, err
	}

	cinemaDomain := []cinemas.Domain{}

	for _, f := range rec {
		cinemaDomain = append(cinemaDomain, f.ToDomain())
	}

	return cinemaDomain, nil
}

func (cinemaRepo *cinemaRepository) GetByID(id string) (cinemas.Domain, error) {
	var cinema Cinema

	if err := cinemaRepo.conn.First(&cinema, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cinemas.Domain{}, errors.New("cinema not found")
		}
		return cinemas.Domain{}, err
	}

	return cinema.ToDomain(), nil
}

func (cinemaRepo *cinemaRepository) Create(cinemaDomain *cinemas.Domain) (cinemas.Domain, error) {
	rec := FromDomain(cinemaDomain)

	if err := cinemaRepo.conn.Create(&rec).Error; err != nil {
		return cinemas.Domain{}, err
	}

	if err := cinemaRepo.conn.Last(&rec).Error; err != nil {
		return cinemas.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cinemaRepo *cinemaRepository) Update(id string, cinemaDomain *cinemas.Domain) (cinemas.Domain, error) {
	var cinema Cinema
	if err := cinemaRepo.conn.First(&cinema, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cinemas.Domain{}, errors.New("cinema not found")
		}
		return cinemas.Domain{}, err
	}

	if cinemaDomain.Name != "" {
		cinema.Name = cinemaDomain.Name
	}

	if cinemaDomain.Location != "" {
		cinema.Location = cinemaDomain.Location
	}

	if cinemaDomain.CityID != 0 {
		cinema.CityID = cinemaDomain.CityID
	}

	if err := cinemaRepo.conn.Save(&cinema).Error; err != nil {
		return cinemas.Domain{}, err
	}

	return cinema.ToDomain(), nil
}


func (cinemaRepo *cinemaRepository) Delete(id string) (bool, error) {
	var cinema Cinema
	if err := cinemaRepo.conn.First(&cinema, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("cinema not found")
		}
		return false, err
	}

	if err := cinemaRepo.conn.Delete(&cinema).Error; err != nil {
		return false, err
	}

	return true, nil
}
