package request

import "kemahin/businesses/tickets"

type Tickets struct {
	EventID int `json:"event_id"`
	UserID  int `json:"user_id"`
	OrderID int `json:"order_id"`
}

func (req *Tickets) ToDomain() *tickets.Domain {
	return &tickets.Domain{
		EventID: req.EventID,
		UserID:  req.UserID,
		OrderID: req.OrderID,
	}
}
