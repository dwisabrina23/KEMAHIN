package request

import "kemahin/businesses/tickets"

type Tickets struct {
	EventID int `json:"id_event"`
	UserID  int `json:"id_user"`
	OrderID int `json:"id_order"`
}

func (req *Tickets) ToDomain() *tickets.Domain {
	return &tickets.Domain{
		EventID: req.EventID,
		UserID:  req.UserID,
		OrderID: req.OrderID,
	}
}
