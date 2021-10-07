package request

import "kemahin/businesses/organizers"

type Organizer struct {
	Id       int    `json: "id"`
	Username string `json: "username"`
	Password string `json: "password"`
	Name     string `json: "name"`
	Phone    string `json: "phone"`
}

type OrgLogin struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func (req *Organizer) ToDomain() *organizers.Domain {
	return &organizers.Domain{
		Id:       req.Id,
		Username: req.Username,
		Pasword:  req.Password,
		Name:     req.Name,
		Phone:    req.Phone,
	}
}
