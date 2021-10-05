package users

import (
	middleware "kemahin/app/middlewares"
	"kemahin/businesses"
	"kemahin/helpers/encrypt"
	"strings"
	// _cacheDomain "kemahin/businesses/cache"
	// "time"
)

type serviceUsers struct {
	repository Repsitory
	jwtAuth    *middleware.ConfigJWT
}

func NewService(repoUSer Repsitory, jtwauth *middleware.ConfigJWT) Service {
	return &serviceUsers{
		repository: repoUSer,
		jwtAuth:    jtwauth,
	}
}

func (su *serviceUsers) Register(data *Domain) (Domain, error) {

	existedUser, err := su.repository.GetByID(data.Id)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}

	if existedUser != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}

	data.Pasword = encrypt.HashAndSalt([]byte(data.Pasword))
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	result, err := su.repository.Register(data)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (su *serviceUsers) Login(nim string, password string) (string, error) {
	if strings.TrimSpace(nim) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrUsernamePasswordInvalid
	}
	userDomain, err := su.repository.GetByNIM(nim)
	if err != nil {
		return "", businesses.ErrUsernamePasswordInvalid
	}

	if !encrypt.ValidateHash(password, userDomain.Pasword) {
		return "", businesses.ErrUsernamePasswordInvalid
	}

	token := su.jwtAuth.GenerateToken(userDomain.Id)
	return token, nil
}

func (su *serviceUsers) Update(data *Domain) (*Domain, error) {
	existedUser, err := su.repository.GetByID(data.Id)
	if err != nil {
		return &Domain{}, err
	}
	data.Id = existedUser.Id

	res, err := su.repository.Update(data)
	if err != nil {
		return &Domain{}, err
	}

	return &res, nil
}

func (su *serviceUsers) GetByID(id int) (Domain, error) {
	resp, err := su.repository.GetByID(id)
	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}
