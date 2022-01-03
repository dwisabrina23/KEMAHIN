package orders

import (
	"kemahin/businesses"
	"kemahin/businesses/events"
	"kemahin/businesses/users"
)

type ServiceOrder struct {
	orderRepository Repository
	eventRepository events.Repository
	userRepository  users.Repository
}

func NewOrderService(orderRepo Repository, eventRepo events.Repository, userRepo users.Repository) Service {
	return &ServiceOrder{
		orderRepository: orderRepo,
		eventRepository: eventRepo,
		userRepository:  userRepo,
	}
}

func (serv *ServiceOrder) Create(idUser int, orderData *Domain) (Domain, error) {
	orderData.UserID = idUser
	userData, err := serv.userRepository.GetByID(idUser)

	eventData, err := serv.eventRepository.GetByID(orderData.EventID)
	if err != nil {
		return Domain{}, businesses.ErrInvalidEventID
	}

	orderData.PaymentName, err = serv.orderRepository.GetPaymentByID(orderData.PaymentID)
	if err != nil {
		return Domain{}, businesses.ErrInvalidPaymentMeth
	}

	if orderData.Qty > eventData.Quota {
		return Domain{}, businesses.ErrOutOfStock
	}

	orderData.UserNIM = userData.NIM
	orderData.Price = eventData.Price * orderData.Qty
	orderData.EventID = eventData.Id
	orderData.EventName = eventData.Judul

	res, err := serv.orderRepository.Create(orderData, eventData)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	return res, nil
}
func (serv *ServiceOrder) GetByUserID(idUSer int) ([]Domain, error) {
	res, err := serv.orderRepository.GetByUserID(idUSer)
	if err != nil {
		return []Domain{}, businesses.ErrOrderNotFound
	}
	return res, nil
}

func (serv *ServiceOrder) ValidateOrder(idOrder int) (Domain, error) {
	resp, err := serv.orderRepository.ValidateOrder(idOrder)
	if err != nil {
		return Domain{}, err
	}

	return resp, nil
}
