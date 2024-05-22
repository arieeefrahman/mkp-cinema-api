package studios


type StudioUseCase struct {
	studioRepo Repository
}

func NewStudioUseCase(studioRepo Repository) Usecase {
	return &StudioUseCase{
		studioRepo: studioRepo,
	}
}

func (studioUc *StudioUseCase) GetAll() ([]Domain, error) {
	return studioUc.studioRepo.GetAll()
}

func (studioUc *StudioUseCase) GetByID(id string) (Domain, error) {
	return studioUc.studioRepo.GetByID(id)
}

func (studioUc *StudioUseCase) Create(studioDomain *Domain) (Domain, error) {
	return studioUc.studioRepo.Create(studioDomain)
}

func (studioUc *StudioUseCase) Update(id string, studioDomain *Domain) (Domain, error) {
	return studioUc.studioRepo.Update(id, studioDomain)
}

func (studioUc *StudioUseCase) Delete(id string) (bool, error) {
	return studioUc.studioRepo.Delete(id)
}
