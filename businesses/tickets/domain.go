package tickets

import "time"

type Domain struct {
	Id        int
	Code      string
	EventID   int
	Prefix    string
	UserID    int
	OrderID   int
	UserEmail string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Create(idOrder int, ticketData *Domain) (Domain, error)
	GetByUserId(idUser int) ([]Domain, error)
}

type Repsitory interface {
	Create(ticketData *Domain) (Domain, error)
	GetByUserId(idUser int) ([]Domain, error)
	GetByID(id int) (Domain, error)
}
