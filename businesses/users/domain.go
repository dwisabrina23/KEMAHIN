package users

import (
	"time"
)

type Domain struct {
	Id        int
	NIM       string
	Pasword   string
	Name      string
	Prodi     string
	Phone     string
	Email     string
	Role      int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Register(data *Domain) (Domain, error)
	Login(nim string, password string) (string, error)
	Update(Data Domain) (Domain, error)
	GetByID(id int) (Domain, error)
}

type Repsitory interface {
	GetByNIM(nim string) (Domain, error)
	Register(data *Domain) (Domain, error)
	Update(Data Domain) (Domain, error)
	GetByID(id int) (Domain, error)
}
