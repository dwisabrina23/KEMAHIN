package events

import (
	"kemahin/businesses"
	"strings"
	"time"
)

type serviceEvents struct {
	repository Repository
}

func NewService(repoEvents Repository) Service {
	return &serviceEvents{
		repository: repoEvents,
	}
}

func (se *serviceEvents) Register(data *Domain) (Domain, error) {

	existedEvent, err := se.repository.GetByID(int(data.Id))
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}

	if existedEvent != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}

	dataEvent, err := se.repository.Register(data)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	return dataEvent, nil
}

func (se *serviceEvents) Update(id int, data *Domain) (Domain, error) {
	existedUser, err := se.repository.GetByID(int(data.Id))
	if err != nil {
		return Domain{}, err
	}
	data.Id = existedUser.Id

	dataEventUpdated, err := se.repository.Update(id, data)
	if err != nil {
		return Domain{}, err
	}

	return dataEventUpdated, nil
}

func (se *serviceEvents) Delete(id int) (string, error) {
	res, err := se.repository.Delete(id)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (se *serviceEvents) GetByID(id int) (Domain, error) {
	resp, err := se.repository.GetByID(id)
	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}

func (se *serviceEvents) UpcomingEvent(date time.Time) ([]Domain, error) {
	res, err := se.repository.UpcomingEvent(date)
	if err != nil {
		return []Domain{}, err
	}
	return res, nil
}

func (se *serviceEvents) GetByJudul(judul string) (Domain, error) {
	res, err := se.repository.GetByJudul(judul)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}
