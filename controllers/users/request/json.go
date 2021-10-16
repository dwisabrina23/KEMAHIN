package request

import "kemahin/businesses/users"

type Users struct {
	Id       int    `json: "id"`
	NIM      string `json: "nim" valid:"required,stringlength(9)"`
	Password string `json: "password"`
	Name     string `json: "name"`
	Prodi    string `json: "prodi"`
	Phone    string `json: "phone"`
	Email    string `json: "email"`
	Role     int    `json: "role"`
}

type UserLogin struct {
	NIM      string `json: "nim" valid:"required,stringlength(9)"`
	Password string `json: "password"`
}

type UserUpdate struct {
	Id       int    `json:"id"`
	Password string `json: "password"`
	Phone    string `json: "phone"`
	Email    string `json: "email"`
}

func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Id:      req.Id,
		NIM:     req.NIM,
		Pasword: req.Password,
		Name:    req.Name,
		Prodi:   req.Prodi,
		Phone:   req.Phone,
		Email:   req.Email,
		Role:    req.Role,
	}
}

func (req *UserUpdate) ToDomain() users.Domain {
	return users.Domain{
		Id:      req.Id,
		Pasword: req.Password,
		Phone:   req.Phone,
		Email:   req.Email,
	}
}
