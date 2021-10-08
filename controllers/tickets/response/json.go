package response

import (
	"kemahin/businesses/tickets"
	"time"
)

type Tickets struct {
	Id        int       `json:"id"`
	Code      string    `json:"ticket_code"`
	EventID   int       `json:"event_id"`
	Prefix    string    `json:"prefix"`
	UserID    int       `json:"user_id"`
	OrderID   int       `json:"order_id"`
	UserEmail string    `json:"user_email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain tickets.Domain) Tickets {
	return Tickets{
		Id:        domain.Id,
		Code:      domain.Code,
		EventID:   domain.EventID,
		Prefix:    domain.Prefix,
		UserID:    domain.UserID,
		OrderID:   domain.OrderID,
		UserEmail: domain.UserEmail,
	}
}

func FromDomainArray(domain []tickets.Domain) []Tickets {
	res := []Tickets{}
	for _, val := range domain {
		res = append(res, FromDomain(val))
	}
	return res
}
