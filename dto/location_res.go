package dto

import "GOLANG/models"

type LocationRes struct {
	ID      int    `json:"id"`
	Place   string `json:"place"`
	Address string `json:"address"`
}

func (ar *LocationRes) FromLocation(a *models.Location) *LocationRes {
	return &LocationRes{
		ID:      a.ID,
		Place:   a.Place,
		Address: a.Address,
	}
}
