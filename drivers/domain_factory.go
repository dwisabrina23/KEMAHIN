package drivers

import (
	userDomain "kemahin/businesses/users"
	userDB "kemahin/drivers/databases/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repsitory {
	return userDB.NewMySqlRepository(conn)
}
