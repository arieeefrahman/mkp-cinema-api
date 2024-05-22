package users

import (
	"mkp-cinema-api/app/middlewares"
)

type UserUsecase struct {
	userRepository Repository
	jwtAuth        *middlewares.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &UserUsecase{
		userRepository: ur,
		jwtAuth:        jwtAuth,
	}
}

func (uu *UserUsecase) Register(userDomain *Domain) (Domain, error) {
	return uu.userRepository.Register(userDomain)
}

func (uu *UserUsecase) Login(userDomain *LoginDomain) (map[string]string, error) {
	tokenPair := make(map[string]string)

	user, err := uu.userRepository.GetByUsername(userDomain)
	if err != nil {
		return tokenPair, err
	}

	token, err := uu.jwtAuth.GenerateToken(user.ID)
	if err != nil {
		return tokenPair, err
	}

	tokenPair["access_token"] = token

	return tokenPair, nil
}
