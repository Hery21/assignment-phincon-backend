package dto

import (
	"GOLANG/models"
	"time"
)

type LogsRes struct {
	ID        int             `json:"id"`
	User      models.User     `json:"user_id"`
	Location  models.Location `json:"location_id"`
	CreatedAt time.Time       `json:"created_at"`
}

func (lr *LogsRes) FromLogs(a *models.Attendance) *LogsRes {
	return &LogsRes{
		ID:        a.ID,
		User:      a.User,
		Location:  a.Location,
		CreatedAt: a.CreatedAt,
	}
}
