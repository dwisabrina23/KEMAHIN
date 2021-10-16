package events

import (
	"time"
)

type Domain struct {
	Id               int
	Judul            string
	Prefix           string
	Poster           string
	Desc             string
	StartDate        time.Time
	EndDate          time.Time
	BatasPendaftaran time.Time
	Place            string
	Quota            int
	Status           int
	Price            int
	CP               string
	Organizer        int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Service interface {
	Register(data *Domain) (Domain, error)
	Update(id int, data *Domain) (Domain, error)
	Delete(id int) (string, error)
	GetByID(id int) (*Domain, error)
	GetByJudul(judul string) ([]Domain, error)
	UpcomingEvent(date time.Time) ([]Domain, error)
}

type Repository interface {
	Register(data *Domain) (Domain, error)
	Update(id int, data *Domain) (Domain, error)
	Delete(id int) (string, error)
	GetByID(id int) (*Domain, error)
	UpcomingEvent(date time.Time) ([]Domain, error)
	GetByJudul(judul string) ([]Domain, error)
}
