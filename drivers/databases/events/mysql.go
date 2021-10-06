package events

import (
	"errors"
	"fmt"
	"kemahin/businesses/events"
	"time"

	"gorm.io/gorm"
)

type mysqlEventsRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) events.Repository {
	return &mysqlEventsRepository{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlEventsRepository) Register(data *events.Domain) (events.Domain, error) {
	recEvent := FromDomain(*data)
	if err := mysqlRepo.Conn.Create(&recEvent).Error; err != nil {
		return events.Domain{}, err
	}
	return recEvent.ToDomain(), nil
}

func (mysqlRepo *mysqlEventsRepository) Update(data *events.Domain) (events.Domain, error) {
	event := FromDomain(*data)
	err := mysqlRepo.Conn.First(&event, data.Id).Updates(*event).Error
	if err != nil {
		return events.Domain{}, err
	}

	return event.ToDomain(), nil
}

func (mysqlRepo *mysqlEventsRepository) Delete(id int) (string, error) {
	rec := Events{}
	nameEvent := rec.Judul
	err := mysqlRepo.Conn.Delete(&rec, "id = ?", id).Error
	if err != nil {
		return "", err
	}
	message := fmt.Sprintf("event %s success to deleted", nameEvent)
	return message, nil
}

func (mysqlRepo *mysqlEventsRepository) GetByID(id int) (events.Domain, error) {
	rec := Events{}
	err := mysqlRepo.Conn.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return events.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (mysqlRepo *mysqlEventsRepository) UpcomingEvent(today time.Time) ([]events.Domain, error) {
	rec := []Events{}
	today = time.Now()
	err := mysqlRepo.Conn.Find(&rec, "start_date > ?", today).Error
	if err != nil {
		return []events.Domain{}, err
	}
	if len(rec) == 0 {
		err = errors.New("theres no upcoming events")
		return []events.Domain{}, nil
	}
	domain := ToArrayOfDomain(rec)
	return domain, nil
}

func (mysqlRepo *mysqlEventsRepository) GetByJudul(judul string) (events.Domain, error) {
	recEvent := Events{}
	err := mysqlRepo.Conn.Find(&recEvent, "judul LIKE %?%", judul).Error
	if err != nil {
		return events.Domain{}, nil
	}

	return recEvent.ToDomain(), nil
}
