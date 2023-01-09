package dto

type RegisterReq struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
	KTPID    string `json:"ktp_id"`
}
