package cinemas

type CinemaUseCase struct {
	cinemaRepo Repository
}

func NewCinemaUseCase(cinemaRepo Repository) Usecase {
	return &CinemaUseCase{
		cinemaRepo: cinemaRepo,
	}
}

func (cu *CinemaUseCase) GetAll() ([]Domain, error) {
	return cu.cinemaRepo.GetAll()
}

func (cu *CinemaUseCase) GetByID(id string) (Domain, error) {
	return cu.cinemaRepo.GetByID(id)
}

func (cu *CinemaUseCase) Create(cinemaDomain *Domain) (Domain, error) {
	return cu.cinemaRepo.Create(cinemaDomain)
}

func (cu *CinemaUseCase) Update(id string, cinemaDomain *Domain) (Domain, error) {
	return cu.cinemaRepo.Update(id, cinemaDomain)
}

func (cu *CinemaUseCase) Delete(id string) (bool, error) {
	return cu.cinemaRepo.Delete(id)
}
