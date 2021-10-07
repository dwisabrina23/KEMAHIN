package response

import (
	"kemahin/businesses/orders"
	"time"
)

type Orders struct {
	Id          int       `gorm:primaryKey`
	UserID      int       `json:"user_id"`
	UserNIM     string    `json:"user_nim"`
	EventID     int       `json:"event_id"`
	EventName   string    `json:"event_name"`
	PaymentID   int       `json:"payment_id"`
	PaymentName string    `json:"payment_name"`
	Price       int       `json:"price"`
	Status      int       `json:"status"`
	Qty         int       `json:"qty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain orders.Domain) Orders {
	return Orders{
		Id:          domain.Id,
		UserID:      domain.UserID,
		UserNIM:     domain.UserNIM,
		EventID:     domain.EventID,
		EventName:   domain.EventName,
		PaymentID:   domain.PaymentID,
		PaymentName: domain.PaymentName,
		Price:       domain.Price,
		Status:      domain.Status,
		Qty:         domain.Qty,
	}
}

func FromDomainArray(domain []orders.Domain) []Orders {
	res := []Orders{}
	for _, val := range domain {
		res = append(res, FromDomain(val))
	}
	return res
}
