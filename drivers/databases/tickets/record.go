package tickets

import (
	"kemahin/businesses/tickets"
	"kemahin/drivers/databases/events"
	"kemahin/drivers/databases/orders"
	"kemahin/drivers/databases/users"

	"gorm.io/gorm"
)

type Tickets struct {
	gorm.Model
	Id      int           `json:"id" gorm:primaryKey`
	Code    string        `json:"code"`
	EventID int           `json:"id_event"`
	Event   events.Events `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	UserID  int           `json:"id_user"`
	User    users.Users   `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;"`
	OrderID int           `json:"id_order"`
	Order   orders.Orders `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:CASCADE;"`
}

func (rec *Tickets) ToDomain() tickets.Domain {
	return tickets.Domain{
		Id:        rec.Id,
		Code:      rec.Code,
		EventID:   rec.EventID,
		Prefix:    rec.Event.Prefix,
		UserID:    rec.UserID,
		UserEmail: rec.User.Email,
		OrderID:   rec.OrderID,
	}
}

func FromDomain(domain tickets.Domain) *Tickets {
	return &Tickets{
		Id:      domain.Id,
		Code:    domain.Code,
		EventID: domain.EventID,
		Event: events.Events{
			Model: gorm.Model{
				ID: uint(domain.EventID),
			},
			Prefix: domain.Prefix,
		},
		UserID: domain.UserID,
		User: users.Users{
			Model: gorm.Model{
				ID: uint(domain.UserID),
			},
			Email: domain.UserEmail,
		},
		OrderID: domain.OrderID,
	}
}

func ToDomainArray(rec []Tickets) []tickets.Domain {
	domain := []tickets.Domain{}
	for _, val := range rec {
		domain = append(domain, val.ToDomain())
	}
	return domain
}
