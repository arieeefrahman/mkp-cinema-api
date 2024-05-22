package cities

import (
	"errors"
	"mkp-cinema-api/businesses/cities"

	"gorm.io/gorm"
)

type cityRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) cities.Repository {
	return &cityRepository{
		conn: conn,
	}
}

func (cityRepo *cityRepository) GetAll() ([]cities.Domain, error) {
	var rec []City

	if err := cityRepo.conn.Find(&rec).Error; err != nil {
		return nil, err
	}

	cityDomain := []cities.Domain{}

	for _, f := range rec {
		cityDomain = append(cityDomain, f.ToDomain())
	}

	return cityDomain, nil
}

func (cityRepo *cityRepository) GetByID(id string) (cities.Domain, error) {
	var city City

	if err := cityRepo.conn.First(&city, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cities.Domain{}, errors.New("city not found")
		}
		return cities.Domain{}, err
	}

	return city.ToDomain(), nil
}

func (cityRepo *cityRepository) Create(cityDomain *cities.Domain) (cities.Domain, error) {
	rec := FromDomain(cityDomain)

	if err := cityRepo.conn.Create(&rec).Error; err != nil {
		return cities.Domain{}, err
	}

	if err := cityRepo.conn.Last(&rec).Error; err != nil {
		return cities.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (cityRepo *cityRepository) Update(id string, cityDomain *cities.Domain) (cities.Domain, error) {
	var city City
	if err := cityRepo.conn.First(&city, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return cities.Domain{}, errors.New("city not found")
		}
		return cities.Domain{}, err
	}

	if cityDomain.Name != "" {
		city.Name = cityDomain.Name
	}

	if err := cityRepo.conn.Save(&city).Error; err != nil {
		return cities.Domain{}, err
	}

	return city.ToDomain(), nil
}


func (cityRepo *cityRepository) Delete(id string) (bool, error) {
	var city City
	if err := cityRepo.conn.First(&city, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("city not found")
		}
		return false, err
	}

	if err := cityRepo.conn.Delete(&city).Error; err != nil {
		return false, err
	}

	return true, nil
}
