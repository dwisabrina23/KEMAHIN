package organizers

import (
	middleware "kemahin/app/middlewares"
	"kemahin/businesses"
	"kemahin/helpers/encrypt"
	"strings"
)

type serviceOrganizer struct {
	repository Repository
	jwtAuth    *middleware.ConfigJWT
}

func NewService(repoOrg Repository, jtwauth *middleware.ConfigJWT) Service {
	return &serviceOrganizer{
		repository: repoOrg,
		jwtAuth:    jtwauth,
	}
}

func (serv *serviceOrganizer) Register(data *Domain) (Domain, error) {

	existedUser, err := serv.repository.GetByUsername(data.Username)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}

	if existedUser != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}

	data.Pasword = encrypt.HashAndSalt([]byte(data.Pasword))
	result, err := serv.repository.Register(data)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (serv *serviceOrganizer) Login(username string, password string) (string, error) {
	if strings.TrimSpace(username) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrUsernamePasswordInvalid
	}
	userDomain, err := serv.repository.GetByUsername(username)
	if err != nil {
		return "", businesses.ErrUsernamePasswordInvalid
	}

	if !encrypt.ValidateHash(password, userDomain.Pasword) {
		return "", businesses.ErrUsernamePasswordInvalid
	}

	token := serv.jwtAuth.GenerateToken(userDomain.Id)
	return token, nil
}

func (serv *serviceOrganizer) GetByID(id int) (Domain, error) {
	resp, err := serv.repository.GetByID(id)
	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}
