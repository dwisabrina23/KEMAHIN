package organizers

import (
	"kemahin/businesses/organizers"

	"gorm.io/gorm"
)

type MySqlOrganizerRepository struct {
	Conn *gorm.DB
}

func NewMySqlRepository(conn *gorm.DB) organizers.Repository {
	return &MySqlOrganizerRepository{
		Conn: conn,
	}
}

func (sqlRepo *MySqlOrganizerRepository) GetByID(id int) (organizers.Domain, error) {
	rec := Organizer{}
	err := sqlRepo.Conn.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return organizers.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (sqlRepo *MySqlOrganizerRepository) GetByUsername(username string) (organizers.Domain, error) {
	rec := Organizer{}
	err := sqlRepo.Conn.Where("username = ?", username).First(&rec).Error
	if err != nil {
		return organizers.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (sqlRepo *MySqlOrganizerRepository) Register(data *organizers.Domain) (organizers.Domain, error) {
	user := fromDomain(*data)
	if err := sqlRepo.Conn.Create(&user).Error; err != nil {
		return organizers.Domain{}, err
	}
	return user.ToDomain(), nil
}
