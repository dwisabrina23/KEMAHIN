package orders

import (
	"context"
	"time"
)

type Domain struct {
	Id         int
	IdPayment  int
	TotalPrice int
	Status     int
	IdBuktiTF  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	GetByID(c context.Context, id int) (Domain, error)
	Store(c context.Context, data *Domain) error
	Update(c context.Context, data *Domain) (Domain, error)
	DeleteById(c context.Context, id int) error
	ValidateBuktiTF(c context.Context, idTF int) error
}

type Repository interface {
	Store(c context.Context, data *Domain) error
	Update(c context.Context, data *Domain) (Domain, error)
	DeleteById(c context.Context, id int) error
	ValidateBuktiTF(c context.Context, idTF int) error
}
