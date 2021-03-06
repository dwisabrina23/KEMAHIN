package events

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"kemahin/businesses/events"
	"time"
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
	// result := mysqlRepo.Conn.Preload("Organizers").Create(&recEvent)
	result := mysqlRepo.Conn.Create(&recEvent)
	if result.Error != nil {
		return events.Domain{}, result.Error
	}
	return recEvent.ToDomain(), nil
}

func (mysqlRepo *mysqlEventsRepository) Update(id int, data *events.Domain) (events.Domain, error) {
	recEvent := FromDomain(*data)

	result := mysqlRepo.Conn.Save(&recEvent)
	if result.Error != nil {
		return events.Domain{}, result.Error
	}

	//update tabel join (event org)
	err := mysqlRepo.Conn.Preload("Organizers").First(&recEvent, recEvent.Id).Error
	if err != nil {
		return events.Domain{}, err
	}

	return recEvent.ToDomain(), nil
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

func (mysqlRepo *mysqlEventsRepository) GetByID(id int) (*events.Domain, error) {
	rec := Events{}
	que := `SELECT * FROM events WHERE events.id = ?`
	err := mysqlRepo.Conn.Raw(que, id).Scan(&rec).Error
	if err != nil {
		return nil, err
	}

	return DetailToDomain(rec), nil
}

// func (mysqlRepo *mysqlEventsRepository) GetByID(id int) (*events.Domain, error) {
// 	rec := Events{}
// 	err := mysqlRepo.Conn.Where("id = ?").First(&rec).Error
// 	// err := mysqlRepo.Conn.Find(&rec, "id = ?", id).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return DetailToDomain(rec), nil
// }

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

func (mysqlRepo *mysqlEventsRepository) GetByJudul(judul string) ([]events.Domain, error) {
	recEvent := []Events{}
	err := mysqlRepo.Conn.Find(&recEvent, "judul LIKE ?", "%"+judul+"%").Error
	if err != nil {
		return []events.Domain{}, nil
	}

	if len(recEvent) == 0 {
		err = errors.New("event nit found")
		return []events.Domain{}, nil
	}

	return ToArrayOfDomain(recEvent), nil
}
