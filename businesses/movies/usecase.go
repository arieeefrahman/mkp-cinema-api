package movies

type MovieUseCase struct {
	movieRepo Repository
}

func NewMovieUsecase(m Repository) Usecase {
	return &MovieUseCase{
		movieRepo: m,
	}
}


func (mu *MovieUseCase) GetAll() ([]Domain, error) {
	return mu.movieRepo.GetAll()
}

func (mu *MovieUseCase) GetByID(id string) (Domain, error) {
	return mu.movieRepo.GetByID(id)
}

func (mu *MovieUseCase) Create(movieDomain *Domain) (Domain, error) {

	return mu.movieRepo.Create(movieDomain)
}

func (mu *MovieUseCase) Update(id string, movieDomain *Domain) (Domain, error) {
	
	return mu.movieRepo.Update(id, movieDomain)
}

func (mu *MovieUseCase) Delete(id string) (bool, error) {
	return mu.movieRepo.Delete(id)
}
