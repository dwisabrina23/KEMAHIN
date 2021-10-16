package response

import (
	"kemahin/businesses/events"
	// organizers "kemahin/controllers/organizer/response"
	"time"
)

type Events struct {
	Id               int       `json: "id"`
	Judul            string    `json: "judul"`
	Prefix           string    `json:"prefix"`
	Poster           string    `json: "poster"`
	Desc             string    `json: "desc"`
	StartDate        time.Time `json: "start" time_format:"2006-01-02 15:04"`
	EndDate          time.Time `json: "end" time_format:"2006-01-02 15:04"`
	BatasPendaftaran time.Time `json: "batas" time_format:"2006-01-02 15:04"`
	Place            string    `json: "place"`
	Quota            int       `json: "quota"`
	Status           int       `json: "status"`
	Price            int       `json: "price"`
	CP               string    `json: "cp"`
	Organizer        int       `json: "organizer"`
	CreatedAt        time.Time `json: "created_at"`
	UpdatedAt        time.Time `json: "updated_at"`
}

func FromDomain(domain events.Domain) Events {
	return Events{
		Id:               domain.Id,
		Judul:            domain.Judul,
		Prefix:           domain.Prefix,
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
		Organizer:        domain.Organizer,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}

func FromDomainDetail(domain *events.Domain) *Events {
	return &Events{
		Id:               domain.Id,
		Judul:            domain.Judul,
		Prefix:           domain.Prefix,
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
		Organizer:        domain.Organizer,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}
