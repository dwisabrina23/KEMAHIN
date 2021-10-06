package drivers

import (
	userDomain "kemahin/businesses/users"
	userDB "kemahin/drivers/databases/users"

	"gorm.io/gorm"
	eventDomain "kemahin/businesses/events"
	eventDB "kemahin/drivers/databases/events"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repsitory {
	return userDB.NewMySqlRepository(conn)
}

func NewEventRepository(conn *gorm.DB) eventDomain.Repository {
	return eventDB.NewMySQLRepository(conn)
}
