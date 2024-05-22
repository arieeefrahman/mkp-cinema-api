package showtimes

import (
	"errors"
	"mkp-cinema-api/businesses/showtimes"

	"gorm.io/gorm"
)

type showtimeRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) showtimes.Repository {
	return &showtimeRepository{
		conn: conn,
	}
}

func (sr *showtimeRepository) GetAll() ([]showtimes.Domain, error) {
	var rec []Showtime

	if err := sr.conn.Find(&rec).Error; err != nil {
		return nil, err
	}

	showtimeDomain := []showtimes.Domain{}

	for _, s := range rec {
		showtimeDomain = append(showtimeDomain, s.ToDomain())
	}

	return showtimeDomain, nil
}

func (sr *showtimeRepository) GetByID(id string) (showtimes.Domain, error) {
	var showtime Showtime

	if err := sr.conn.First(&showtime, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return showtimes.Domain{}, errors.New("showtime not found")
		}
		return showtimes.Domain{}, err
	}

	return showtime.ToDomain(), nil
}

func (sr *showtimeRepository) Create(showtimeDomain *showtimes.Domain) (showtimes.Domain, error) {
	rec := FromDomain(showtimeDomain)

	if err := sr.conn.Create(&rec).Error; err != nil {
		return showtimes.Domain{}, err
	}

	if err := sr.conn.Last(&rec).Error; err != nil {
		return showtimes.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (sr *showtimeRepository) Update(id string, showtimeDomain *showtimes.Domain) (showtimes.Domain, error) {
	var showtime Showtime
	if err := sr.conn.First(&showtime, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return showtimes.Domain{}, errors.New("showtime not found")
		}
		return showtimes.Domain{}, err
	}

	if !showtimeDomain.Date.IsZero() {
		showtime.Date = showtimeDomain.Date
	}
	if !showtimeDomain.StartTime.IsZero() {
		showtime.StartTime = showtimeDomain.StartTime
	}
	if !showtimeDomain.EndTime.IsZero() {
		showtime.EndTime = showtimeDomain.EndTime
	}

	if err := sr.conn.Save(&showtime).Error; err != nil {
		return showtimes.Domain{}, err
	}

	return showtime.ToDomain(), nil
}


func (sr *showtimeRepository) Delete(id string) (bool, error) {
	var showtime Showtime
	if err := sr.conn.First(&showtime, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("showtime not found")
		}
		return false, err
	}

	if err := sr.conn.Delete(&showtime).Error; err != nil {
		return false, err
	}

	return true, nil
}
