package response

import (
	"kemahin/businesses/events"
	"time"
)

type Events struct {
	Id               uint      `json: "id"`
	Judul            string    `json: "judul"`
	Poster           string    `json: "poster"`
	Desc             string    `json: "desc"`
	StartDate        time.Time `json: "start_date"`
	EndDate          time.Time `json: "end_date"`
	BatasPendaftaran time.Time `json: "batas_pendaftaran"`
	Place            string    `json: "place"`
	Quota            int       `json: "quota"`
	Status           int       `json: "status"`
	Price            int       `json: "price"`
	CP               string    `json: "cp"`
	IDOrganizer      int       `json: "id_org"`
}

func FromDomain(domain events.Domain) Events {
	return Events{
		Id:               domain.Id,
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
