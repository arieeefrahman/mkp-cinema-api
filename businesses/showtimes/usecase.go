package showtimes

type ShowtimeUsecase struct {
	showtimeRepo Repository
}

func NewShowtimeUsecase(showtimeRepo Repository) Usecase {
	return &ShowtimeUsecase{
		showtimeRepo: showtimeRepo,
	}
}

func (showtimeUc *ShowtimeUsecase) GetAll() ([]Domain, error) {
	return showtimeUc.showtimeRepo.GetAll()
}

func (showtimeUc *ShowtimeUsecase) GetByID(id string) (Domain, error) {
	return showtimeUc.showtimeRepo.GetByID(id)
}

func (showtimeUc *ShowtimeUsecase) Create(showtimeDomain *Domain) (Domain, error) {
	return showtimeUc.showtimeRepo.Create(showtimeDomain)
}

func (showtimeUc *ShowtimeUsecase) Update(id string, showtimeDomain *Domain) (Domain, error) {
	return showtimeUc.showtimeRepo.Update(id, showtimeDomain)
}

func (showtimeUc *ShowtimeUsecase) Delete(id string) (bool, error) {
	return showtimeUc.showtimeRepo.Delete(id)
}
