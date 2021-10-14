package users

import (
	"kemahin/businesses/users"
	"time"
	// "gorm.io/gorm"
)

type Users struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	NIM       string    `json:"nim" gorm:"unique"`
	Pasword   string    `json:"password"`
	Name      string    `json:"name"`
	Prodi     string    `json:"prodi"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		Id:        rec.Id,
		NIM:       rec.NIM,
		Pasword:   rec.Pasword,
		Name:      rec.Name,
		Prodi:     rec.Prodi,
		Phone:     rec.Phone,
		Email:     rec.Email,
		RoleID:    rec.RoleID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(domain users.Domain) Users {
	return Users{
		Id:      domain.Id,
		NIM:     domain.NIM,
		Pasword: domain.Pasword,
		Name:    domain.Name,
		Prodi:   domain.Prodi,
		Phone:   domain.Phone,
		Email:   domain.Email,
		RoleID:  domain.RoleID,
	}
}
