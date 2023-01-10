package dto

type ForgotPasswordReq struct {
	KTPID    string `json:"ktp_id"`
	Password string `json:"password"`
}
