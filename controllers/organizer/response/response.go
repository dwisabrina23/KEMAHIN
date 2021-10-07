package response

import (
	"kemahin/businesses/organizers"
	"time"
)

type Organizer struct {
	Id        int       `json:"id"`
	Username  string    `json: "username"`
	Pasword   string    `json: "password"`
	Name      string    `json: "name"`
	Phone     string    `json: "Phone"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

func FromDomain(domain organizers.Domain) Organizer {
	return Organizer{
		Id:        domain.Id,
		Username:  domain.Username,
		Pasword:   domain.Pasword,
		Name:      domain.Name,
		Phone:     domain.Phone,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
