package orders

import (
	"kemahin/businesses/events"
	"time"
)

type Domain struct {
	Id          int
	UserID      int
	UserNIM     string
	EventID     int
	EventName   string
	PaymentID   int
	PaymentName string
	Price       int
	Status      int
	Qty         int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// IdBuktiTF  int

}

type Service interface {
	Create(idUser int, orderData *Domain) (Domain, error)
	GetByUserID(idUSer int) ([]Domain, error)
	ValidateOrder(idOrder int) (Domain, error)
}

type Repository interface {
	Create(orderData *Domain, eventData *events.Domain) (Domain, error)
	GetByUserID(idUSer int) ([]Domain, error)
	ValidateOrder(idOrder int) (Domain, error)
	GetPaymentByID(idPay int) (string, error)
}
