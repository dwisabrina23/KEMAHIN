package orders

import (
	"errors"
	"kemahin/businesses/events"
	"kemahin/businesses/orders"

	"gorm.io/gorm"
)

type mySqlOrdersRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) orders.Repository {
	return &mySqlOrdersRepository{
		Conn: conn,
	}
}
func (mysqlRepo *mySqlOrdersRepository) Create(orderData *orders.Domain, eventData *events.Domain) (orders.Domain, error) {
	rec := FromDomain(*orderData)

	if err := mysqlRepo.Conn.Create(&rec).Error; err != nil {
		return orders.Domain{}, nil
	}

	return rec.ToDomain(), nil
}

func (mysqlRepo *mySqlOrdersRepository) GetByUserID(userId int) ([]orders.Domain, error) {
	rec := []Orders{}
	err := mysqlRepo.Conn.Joins("Events").Joins("Payment").Find(&rec, "user_id", userId).Error
	if len(rec) == 0 {
		err = errors.New("Order data not found")
		return []orders.Domain{}, err
	}

	return ToDomainArray(rec), nil
}

func (mysqlRepo *mySqlOrdersRepository) ValidateOrder(idOrder int) (orders.Domain, error) {
	rec := Orders{}
	err := mysqlRepo.Conn.Joins("Events").Joins("Payment").Where("id = ?", idOrder).Update("status", 1).Error

	if err != nil {
		return orders.Domain{}, err
	}

	return rec.ToDomain(), nil
}

func (mysqlRepo *mySqlOrdersRepository) GetPaymentByID(idPay int) (string, error) {
	paymentRec := Payment{}
	err := mysqlRepo.Conn.First(&paymentRec, idPay).Error
	if err != nil {
		return "", err
	}
	return paymentRec.Name, nil
}

func (mysqlRepo *mySqlOrdersRepository) GetByOrderId(id int) (orders.Domain, error) {
	rec := Orders{}
	err := mysqlRepo.Conn.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return orders.Domain{}, err
	}
	return rec.ToDomain(), nil
}
