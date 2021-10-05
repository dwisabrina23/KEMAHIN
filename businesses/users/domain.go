package users

import (
	"time"
)

type Domain struct {
	Id        int       `json:"id"`
	NIM       string    `json:"nim"`
	Pasword   string    `json:"password"`
	Name      string    `json:"name"`
	Prodi     string    `json:"prodi"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Service interface {
	Register(data *Domain) (Domain, error)
	Login(nim string, password string) (string, error)
	Update(data *Domain) (*Domain, error)
	GetByID(id int) (Domain, error)
}

type Repsitory interface {
	GetByNIM(nim string) (Domain, error)
	Register(data *Domain) (Domain, error)
	Update(data *Domain) (Domain, error)
	GetByID(id int) (Domain, error)
}
