package organizers

import (
	"gorm.io/gorm"
	"kemahin/businesses/organizers"
)

type Organizers struct {
	gorm.Model
	Username string `json:"username"`
	Pasword  string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}

func (rec *Organizers) ToDomain() organizers.Domain {
	return organizers.Domain{
		Id:        int(rec.ID),
		Username:  rec.Username,
		Pasword:   rec.Pasword,
		Name:      rec.Name,
		Phone:     rec.Phone,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(domain organizers.Domain) *Organizers {
	return &Organizers{
		Model: gorm.Model{
			ID:        uint(domain.Id),
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		Username: domain.Username,
		Pasword:  domain.Pasword,
		Name:     domain.Name,
		Phone:    domain.Phone,
	}
}
