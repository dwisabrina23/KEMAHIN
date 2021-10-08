package request

import (
	"kemahin/businesses/events"
	"time"
)

type Events struct {
	Id               int       `json: "id"`
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

func (rec *Events) ToDomain() *events.Domain {
	return &events.Domain{
		Id:               rec.Id,
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
