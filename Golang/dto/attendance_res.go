package dto

import "GOLANG/models"

type AttendanceRes struct {
	ID         int             `json:"id"`
	CheckInOut string          `json:"check_in_out"`
	User       models.User     `json:"user_id"`
	Location   models.Location `json:"location_id"`
}

func (ar *AttendanceRes) FromAttendance(a *models.Attendance) *AttendanceRes {
	return &AttendanceRes{
		ID:         a.ID,
		CheckInOut: a.CheckInOut,
		User:       a.User,
		Location:   a.Location,
	}
}
