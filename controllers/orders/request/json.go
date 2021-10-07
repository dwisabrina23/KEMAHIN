package request

import "kemahin/businesses/orders"

type Orders struct {
	UserID    int `json:"user_id"`
	EventID   int `json:"event_id"`
	PaymentID int `json:"payment_id"`
	Qty       int `json: "qty"`
}

func (req *Orders) ToDomain() *orders.Domain {
	return &orders.Domain{
		UserID:    req.UserID,
		EventID:   req.EventID,
		PaymentID: req.PaymentID,
		Qty:       req.Qty,
	}
}
