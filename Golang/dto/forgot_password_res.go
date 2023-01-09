package dto

import "GOLANG/models"

type ForgotPasswordRes struct {
	KTPID string `json:"ktp_id"`
}

func (fpr *ForgotPasswordRes) FromForgotPassword(f *models.User) *ForgotPasswordRes {
	return &ForgotPasswordRes{
		KTPID: f.KTPID,
	}
}
