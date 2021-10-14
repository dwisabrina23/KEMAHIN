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
	Id      int           `json:"id"`
	Code    string        `json:"code"`
	EventID int           `json:"event_id"`
	Event   events.Events `gorm:"foreignKey:EventID;references:ID"`
	UserID  int           `json:"user_id"`
	User    users.Users   `gorm:"foreignKey:UserID;references:ID"`
	OrderID int           `json:"order_id"`
	Order   orders.Orders `gorm:"foreignKey:OrderID;references:ID"`
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
			Id: domain.UserID,
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
