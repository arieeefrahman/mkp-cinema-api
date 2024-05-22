package cities

type CityUseCase struct {
	cityRepo Repository
}

func NewCityUseCase(cityRepo Repository) Usecase {
	return &CityUseCase{
		cityRepo: cityRepo,
	}
}


func (cityUc *CityUseCase) GetAll() ([]Domain, error) {
	return cityUc.cityRepo.GetAll()
}

func (cityUc *CityUseCase) GetByID(id string) (Domain, error) {
	return cityUc.cityRepo.GetByID(id)
}

func (cityUc *CityUseCase) Create(cityDomain *Domain) (Domain, error) {
	return cityUc.cityRepo.Create(cityDomain)
}

func (cityUc *CityUseCase) Update(id string, cityDomain *Domain) (Domain, error) {
	return cityUc.cityRepo.Update(id, cityDomain)
}

func (cityUc *CityUseCase) Delete(id string) (bool, error) {
	return cityUc.cityRepo.Delete(id)
}
