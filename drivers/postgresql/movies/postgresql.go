package movies

import (
	"errors"
	"mkp-cinema-api/businesses/movies"

	"gorm.io/gorm"
)

type movieRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) movies.Repository {
	return &movieRepository{
		conn: conn,
	}
}

func (mr *movieRepository) GetAll() ([]movies.Domain, error) {
	var rec []Movie

	if err := mr.conn.Find(&rec).Error; err != nil {
		return nil, err
	}

	movieDomain := []movies.Domain{}

	for _, f := range rec {
		movieDomain = append(movieDomain, f.ToDomain())
	}

	return movieDomain, nil
}

func (mr *movieRepository) GetByID(id string) (movies.Domain, error) {
	var movie Movie

	if err := mr.conn.First(&movie, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return movies.Domain{}, errors.New("movie not found")
		}
		return movies.Domain{}, err
	}

	return movie.ToDomain(), nil
}

func (mr *movieRepository) Create(movieDomain *movies.Domain) (movies.Domain, error) {
	rec := FromDomain(movieDomain)

	if err := mr.conn.Create(&rec).Error; err != nil {
		return movies.Domain{}, err
	}

	if err := mr.conn.Last(&rec).Error; err != nil {
		return movies.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (mr *movieRepository) Update(id string, movieDomain *movies.Domain) (movies.Domain, error) {
	var movie Movie
	if err := mr.conn.First(&movie, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return movies.Domain{}, errors.New("movie not found")
		}
		return movies.Domain{}, err
	}

	// Apply updates only to the fields that are provided in the movieDomain
	if movieDomain.Title != "" {
		movie.Title = movieDomain.Title
	}
	if movieDomain.Genre != "" {
		movie.Genre = movieDomain.Genre
	}
	if movieDomain.Duration != 0 {
		movie.Duration = movieDomain.Duration
	}
	if movieDomain.Rating != 0 {
		movie.Rating = movieDomain.Rating
	}
	if !movieDomain.ReleaseDate.IsZero() {
		movie.ReleaseDate = movieDomain.ReleaseDate
	}
	if movieDomain.Description != "" {
		movie.Description = movieDomain.Description
	}

	if err := mr.conn.Save(&movie).Error; err != nil {
		return movies.Domain{}, err
	}

	return movie.ToDomain(), nil
}


func (mr *movieRepository) Delete(id string) (bool, error) {
	var movie Movie
	if err := mr.conn.First(&movie, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("movie not found")
		}
		return false, err
	}

	if err := mr.conn.Delete(&movie).Error; err != nil {
		return false, err
	}

	return true, nil
}
