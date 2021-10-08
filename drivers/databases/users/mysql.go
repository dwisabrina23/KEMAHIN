package users

import (
	"kemahin/businesses/users"

	"gorm.io/gorm"
)

type MySqlUsersRepository struct {
	Conn *gorm.DB
}

func NewMySqlRepository(conn *gorm.DB) users.Repsitory {
	return &MySqlUsersRepository{
		Conn: conn,
	}
}

func (ur *MySqlUsersRepository) Register(userData *users.Domain) (users.Domain, error) {
	user := fromDomain(*userData)
	// user.Pasword = _encrypt.HashAndSalt([]byte(user.Pasword))
	if err := ur.Conn.Create(&user).Error; err != nil {
		return users.Domain{}, err
	}
	return user.toDomain(), nil
}

func (ur *MySqlUsersRepository) Update(id int, data *users.Domain) (users.Domain, error) {
	user := fromDomain(*data)
	err := ur.Conn.First(&user, "id = ?", id).Updates(user).Error
	if err != nil {
		return users.Domain{}, err
	}

	return user.toDomain(), nil
}

func (ur *MySqlUsersRepository) GetByID(id int) (users.Domain, error) {
	rec := Users{}
	err := ur.Conn.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (ur *MySqlUsersRepository) GetByNIM(nim string) (users.Domain, error) {
	rec := Users{}
	err := ur.Conn.Where("nim = ?", nim).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}
