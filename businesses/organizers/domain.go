package organizers

import "time"

type Domain struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Pasword   string    `json:"password"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Service interface {
	Login(nim string, password string) (string, error)
	Register(data *Domain) (Domain, error)
	GetByID(id int) (Domain, error)
}

type Repository interface {
	GetByID(id int) (Domain, error)
	Register(data *Domain) (Domain, error)
	GetByUsername(name string) (Domain, error)
}
