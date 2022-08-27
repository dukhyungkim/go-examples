package entity

import (
	"dvdrental/graph/model"
	"dvdrental/repository"
	"fmt"
)

type Film struct {
	model.Film
}

func NewFilm(f *repository.Film) *Film {
	var desc *string
	if f.Description.Valid {
		desc = &f.Description.String
	}

	var length *int
	if f.Length.Valid {
		value := int(f.Length.Int16)
		length = &value
	}

	return &Film{
		model.Film{
			ID:              fmt.Sprint(f.FilmID),
			Title:           f.Title,
			Description:     desc,
			ReleaseYear:     int(f.ReleaseYear.(int64)),
			LanguageID:      fmt.Sprint(f.LanguageID),
			RentalDuration:  int(f.RentalDuration),
			RentalRate:      f.RentalRate,
			Length:          length,
			ReplacementCost: f.ReplacementCost,
			Rating:          string(f.Rating.([]uint8)),
			LastUpdate:      f.LastUpdate,
			SpecialFeatures: f.SpecialFeatures,
			Fulltext:        string(f.Fulltext.([]uint8)),
		},
	}
}

func (f *Film) ToModel() *model.Film {
	return &f.Film
}
