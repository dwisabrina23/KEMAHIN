package response

import (
	"kemahin/businesses/users"
	"time"
)

type Users struct {
	Id        int       `json:"id"`
	NIM       string    `json: "nim"`
	Pasword   string    `json: "password"`
	Name      string    `json: "name"`
	Prodi     string    `json: "prodi"`
	Phone     string    `json: "phone"`
	Email     string    `json: "email"`
	Role      int       `json: "role"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Id:        domain.Id,
		NIM:       domain.NIM,
		Pasword:   domain.Pasword,
		Name:      domain.Name,
		Prodi:     domain.Prodi,
		Phone:     domain.Phone,
		Email:     domain.Email,
		Role:      domain.Role,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
