package sendgrids

import (
	"kemahin/businesses"
	"kemahin/businesses/tickets"
)

type SendService struct {
	sendRepository   Repository
	ticketRepository tickets.Repsitory
}

func NewSendService(ticketRepo tickets.Repsitory, sendRepo Repository) Service {
	return &SendService{
		sendRepository:   sendRepo,
		ticketRepository: ticketRepo,
	}
}

func (serv *SendService) Send(idTicket int, emailData *Domain) (Domain, error) {
	TicketData, err := serv.ticketRepository.GetByID(idTicket)
	if err != nil {
		return Domain{}, businesses.ErrInvalidTicketID
	}
	emailData.TicketID = idTicket
	emailData.TicketCode = TicketData.Code
	emailData.UserId = TicketData.UserID
	emailData.UserEmail = TicketData.UserEmail

	data, errs := serv.sendRepository.Send(emailData)
	if errs != nil {
		return Domain{}, businesses.ErrInternalServer
	}
	return data, nil
}
