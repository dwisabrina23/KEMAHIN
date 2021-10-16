package drivers

import (
	"gorm.io/gorm"

	userDomain "kemahin/businesses/users"
	userDB "kemahin/drivers/databases/users"

	eventDomain "kemahin/businesses/events"
	eventDB "kemahin/drivers/databases/events"

	orgDomain "kemahin/businesses/organizers"
	orgDB "kemahin/drivers/databases/organizers"

	orderDomain "kemahin/businesses/orders"
	orderDB "kemahin/drivers/databases/orders"

	ticketDomain "kemahin/businesses/tickets"
	ticketDB "kemahin/drivers/databases/tickets"

	sendDomain "kemahin/businesses/sendgrids"
	sendDB "kemahin/drivers/thirdparties/sendgrid"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repsitory {
	return userDB.NewMySqlRepository(conn)
}

func NewOrgRepository(conn *gorm.DB) orgDomain.Repository {
	return orgDB.NewMySqlRepository(conn)
}

func NewEventRepository(conn *gorm.DB) eventDomain.Repository {
	return eventDB.NewMySQLRepository(conn)
}

func NewOrderRepository(conn *gorm.DB) orderDomain.Repository {
	return orderDB.NewMySQLRepository(conn)
}

func NewTicketRepository(conn *gorm.DB) ticketDomain.Repsitory {
	return ticketDB.NewMySQLRepository(conn)
}

func NewSendRepository(conn *gorm.DB) sendDomain.Repository {
	return sendDB.NewSendAPI()
}
