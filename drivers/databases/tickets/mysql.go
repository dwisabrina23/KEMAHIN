package tickets

import (
	"errors"
	"kemahin/businesses/tickets"

	"gorm.io/gorm"
)

type mySqlTicketRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) tickets.Repository {
	return &mySqlTicketRepository{
		Conn: conn,
	}
}

func (mysqlRepo *mySqlTicketRepository) Create(ticketData *tickets.Domain) (tickets.Domain, error) {
	rec := FromDomain(*ticketData)

	if err := mysqlRepo.Conn.Create(&rec).Error; err != nil {
		return tickets.Domain{}, err
	}
	return rec.ToDomain(), nil

}

func (mysqlRepo *mySqlTicketRepository) GetByUserId(idUser int) ([]tickets.Domain, error) {
	rec := []Tickets{}
	err := mysqlRepo.Conn.Joins("Events").Joins("Orders").Find(&rec, "user_id", idUser).Error
	if err != nil {
		return []tickets.Domain{}, err
	}
	if len(rec) == 0 {
		err = errors.New("order data not found")
		return []tickets.Domain{}, err
	}

	return ToDomainArray(rec), nil
}

func (mysqlRepo *mySqlTicketRepository) GetByID(id int) (tickets.Domain, error) {
	rec := Tickets{}
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Error
	if err != nil {
		return tickets.Domain{}, err
	}
	return rec.ToDomain(), nil
}
