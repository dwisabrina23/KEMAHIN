package events

import (
	"kemahin/businesses/events"
	"kemahin/drivers/databases/organizers"
	"time"

	"gorm.io/gorm"
)

type Events struct {
	gorm.Model
	Judul            string               `json: "judul"`
	Poster           string               `json: "poster"`
	Desc             string               `json: "desc"`
	StartDate        time.Time            `json: "start_date"`
	EndDate          time.Time            `json: "end_date"`
	BatasPendaftaran time.Time            `json: "batas_pendaftaran"`
	Place            string               `json: "place"`
	Quota            int                  `json: "quota"`
	Status           int                  `json: "status"`
	Price            int                  `json: "price"`
	CP               string               `json: "cp"`
	IDOrganizer      int                  `json: "id_org"`
	Organizer        organizers.Organizer `gorm:"constraint:OnUpdate:CASCADE"`
}

func (rec *Events) ToDomain() events.Domain {
	return events.Domain{
		Judul:            rec.Judul,
		Poster:           rec.Poster,
		Desc:             rec.Desc,
		StartDate:        rec.StartDate,
		EndDate:          rec.EndDate,
		BatasPendaftaran: rec.BatasPendaftaran,
		Place:            rec.Place,
		Quota:            rec.Quota,
		Status:           rec.Status,
		Price:            rec.Price,
		CP:               rec.CP,
		IDOrganizer:      rec.IDOrganizer,
	}
}

func ToArrayOfDomain(rec []Events) []events.Domain {
	domainArray := []events.Domain{}

	for _, val := range rec {
		domainArray = append(domainArray, val.ToDomain())
	}

	return domainArray
}

func FromDomain(domain events.Domain) *Events {
	return &Events{
		Model: gorm.Model{
			ID:        uint(domain.Id),
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		Judul:            domain.Judul,
		Poster:           domain.Poster,
		Desc:             domain.Desc,
		StartDate:        domain.StartDate,
		EndDate:          domain.EndDate,
		BatasPendaftaran: domain.BatasPendaftaran,
		Place:            domain.Place,
		Quota:            domain.Quota,
		Status:           domain.Status,
		Price:            domain.Price,
		CP:               domain.CP,
		IDOrganizer:      domain.IDOrganizer,
	}
}
