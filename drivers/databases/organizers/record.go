package organizers

import (
	"kemahin/businesses/organizers"
	"time"
	// "gorm.io/gorm"
)

type Organizers struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Pasword   string    `json:"password"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (rec *Organizers) ToDomain() organizers.Domain {
	return organizers.Domain{
		Id:        rec.Id,
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
		Id:        domain.Id,
		Username:  domain.Username,
		Pasword:   domain.Pasword,
		Name:      domain.Name,
		Phone:     domain.Phone,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
