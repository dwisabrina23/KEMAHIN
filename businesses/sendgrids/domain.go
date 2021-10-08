package sendgrids

type Domain struct {
	Id         int
	TicketID   int
	TicketCode string
	UserId     int
	UserEmail  string
	QRLink     string
}
type Service interface {
	Send(ticketID int, emailData *Domain) (Domain, error)
}

type Repository interface {
	Send(emailData *Domain) (Domain, error)
}
