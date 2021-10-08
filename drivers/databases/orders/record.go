package orders

import (
	"kemahin/businesses/orders"
	"kemahin/drivers/databases/events"
	"kemahin/drivers/databases/users"

	"gorm.io/gorm"
)

type Payment struct {
	Id   int `gorm:"primaryKey`
	Name string
}
type Orders struct {
	Id        int           `json:"id" gorm:primaryKey`
	UserID    int           `json:"user_id"`
	User      users.Users   `gorm:"foreignKey:UserID;references:ID"`
	EventID   int           `json:"event_id"`
	Event     events.Events `gorm:"foreignKey:EventID;references:ID"`
	PaymentID int           `json:"payment_id"`
	Payment   Payment       `gorm:"foreignKey:PaymentID;references:ID"`
	Price     int           `json:"price"`
	Status    int           `json:"status" gorm:"default:0"`
	Qty       int           `json:"qty"`
}

func (rec *Orders) ToDomain() orders.Domain {
	return orders.Domain{
		Id:          rec.Id,
		UserID:      rec.UserID,
		EventID:     rec.EventID,
		EventName:   rec.Event.Judul,
		PaymentID:   rec.PaymentID,
		PaymentName: rec.Payment.Name,
		Price:       rec.Price,
		Status:      rec.Status,
		Qty:         rec.Qty,
	}
}

func ToDomainArray(rec []Orders) []orders.Domain {
	domain := []orders.Domain{}

	for _, val := range rec {
		domain = append(domain, val.ToDomain())
	}
	return domain
}

func FromDomain(domain orders.Domain) *Orders {
	return &Orders{
		Id:      domain.Id,
		UserID:  domain.UserID,
		EventID: domain.EventID,
		Event: events.Events{
			Model: gorm.Model{
				ID: uint(domain.EventID),
			},
			Judul: domain.EventName,
		},
		Payment: Payment{domain.PaymentID, domain.PaymentName},
		Price:   domain.Price,
		Status:  domain.Status,
		Qty:     domain.Qty,
	}
}
