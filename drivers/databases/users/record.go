package users

import (
	"kemahin/businesses/users"
	"time"
)

type Users struct {
	ID        int
	NIM       string
	Pasword   string
	Name      string
	Prodi     string
	Phone     string
	Email     string
	RoleID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		Id:        rec.ID,
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

func fromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:        userDomain.Id,
		NIM:       userDomain.NIM,
		Pasword:   userDomain.Pasword,
		Name:      userDomain.Name,
		Prodi:     userDomain.Prodi,
		Phone:     userDomain.Phone,
		Email:     userDomain.Email,
		RoleID:    userDomain.RoleID,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}
