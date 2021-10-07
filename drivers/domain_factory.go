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
)

func NewUserRepository(conn *gorm.DB) userDomain.Repsitory {
	return userDB.NewMySqlRepository(conn)
}

func NewEventRepository(conn *gorm.DB) eventDomain.Repository {
	return eventDB.NewMySQLRepository(conn)
}

func NewOrgRepository(conn *gorm.DB) orgDomain.Repository {
	return orgDB.NewMySqlRepository(conn)
}

func NewOrderRepository(conn *gorm.DB) orderDomain.Repository {
	return orderDB.NewMySQLRepository(conn)
}
