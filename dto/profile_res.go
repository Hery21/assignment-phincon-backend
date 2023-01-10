package dto

import "GOLANG/models"

type ProfileRes struct {
	FullName string `json:"full_name"`
	KTPID    string `json:"ktp_id"`
	Role     string `json:"role"`
	Address  string `json:"address"`
}

func (pr *ProfileRes) FromUser(r *models.User) *ProfileRes {
	return &ProfileRes{
		FullName: r.FullName,
		KTPID:    r.KTPID,
		Role:     r.Role,
		Address:  r.Address,
	}
}
