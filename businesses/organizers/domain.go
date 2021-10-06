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
	GetByID(id int) (Domain, error)
	Update(id int, orgData *Domain) (Domain, error)
}

type Repository interface {
	GetByID(id int) (Domain, error)
	Update(id int, orgData *Domain) (Domain, error)
	GetByName(name string) (Domain, error)
}
