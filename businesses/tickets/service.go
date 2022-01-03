package tickets

import (
	"kemahin/businesses"
	"kemahin/businesses/events"
	"kemahin/businesses/orders"
	"kemahin/businesses/users"
	"kemahin/helpers/ticketing/tickets"
)

type ticketService struct {
	ticketRepository Repository
	userRepository   users.Repository
	eventRepository  events.Repository
	orderRepository  orders.Repository
}

func NewTicketService(ticketRepo Repository, userRepo users.Repository, eventRepo events.Repository) Service {
	return &ticketService{
		ticketRepository: ticketRepo,
		userRepository:   userRepo,
		eventRepository:  eventRepo,
	}
}

func (serv *ticketService) Create(idOrder int, ticketData *Domain) (Domain, error) {
	ticketData.OrderID = idOrder
	orderData, err := serv.orderRepository.GetByOrderId(idOrder)
	if err != nil {
		return Domain{}, businesses.ErrInvalidOrderID
	}
	//save user id from user data get by order id
	ticketData.UserID = orderData.UserID
	ticketData.EventID = orderData.EventID

	eventData, err := serv.eventRepository.GetByID(orderData.EventID)
	if err != nil {
		return Domain{}, err
	}

	ticketData.Prefix = eventData.Prefix
	ticketData.Code = tickets.GenerateTicketCode(ticketData.Prefix, idOrder)

	userData, err := serv.userRepository.GetByID(orderData.UserID)
	ticketData.UserEmail = userData.Email

	resp, err := serv.ticketRepository.Create(ticketData)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	return resp, nil

}
func (serv *ticketService) GetByUserId(idUser int) ([]Domain, error) {
	res, err := serv.ticketRepository.GetByUserId(idUser)
	if err != nil {
		return []Domain{}, businesses.ErrOrderNotFound
	}
	return res, nil
}
