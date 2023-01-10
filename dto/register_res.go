package dto

import "GOLANG/models"

type RegisterRes struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
	KTPID    string `json:"ktp_id"`
}

func (rr *RegisterRes) FromRegister(r *models.User) *RegisterRes {
	return &RegisterRes{
		ID:       r.ID,
		Username: r.Username,
		FullName: r.FullName,
		KTPID:    r.KTPID,
	}
}
