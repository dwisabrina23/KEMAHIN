package request

import (
	"kemahin/businesses/events"
	"time"
)

type Events struct {
	// Id               int       `json: "id"`
	Judul            string `json: "judul"`
	Prefix           string `json:"prefix"`
	Poster           string `json: "poster"`
	Desc             string `json: "desc"`
	StartDate        string `json: "start" time_format:"2006-01-02 15:04"`
	EndDate          string `json: "end" time_format:"2006-01-02 15:04"`
	BatasPendaftaran string `json: "batas" time_format:"2006-01-02 15:04"`
	Place            string `json: "place"`
	Quota            int    `json: "quota"`
	Status           int    `json: "status"`
	Price            int    `json: "price"`
	CP               string `json: "cp"`
	Organizer        int    `json: "organizer"`
}

func (rec *Events) ToDomain() *events.Domain {
	startDate, _ := time.Parse("2006-01-02 15:04", rec.StartDate)
	endDate, _ := time.Parse("2006-01-02 15:04", rec.EndDate)
	batasDate, _ := time.Parse("2006-01-02 15:04", rec.BatasPendaftaran)
	return &events.Domain{
		Judul:            rec.Judul,
		Prefix:           rec.Prefix,
		Poster:           rec.Poster,
		Desc:             rec.Desc,
		StartDate:        startDate,
		EndDate:          endDate,
		BatasPendaftaran: batasDate,
		Place:            rec.Place,
		Quota:            rec.Quota,
		Status:           rec.Status,
		Price:            rec.Price,
		CP:               rec.CP,
		Organizer:        rec.Organizer,
	}
}
