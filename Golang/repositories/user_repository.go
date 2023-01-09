package repositories

import (
	"GOLANG/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	MatchingCredential(username string) (*models.User, error)
	ForgotPassword(*models.User) (*models.User, error)
	Location() ([]*models.Location, error)
	CreateAttendance(*models.Attendance) (*models.Attendance, error)
	CreateCheckout(*models.Attendance) (*models.Attendance, error)
	Logs(string) ([]*models.Attendance, error)
	Profile(id int) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) *userRepository {
	return &userRepository{db: c.DB}
}

func (u *userRepository) MatchingCredential(username string) (*models.User, error) {
	var user *models.User

	res := u.db.Where("username = ?", username).First(&user)

	isNotFound := errors.Is(res.Error, gorm.ErrRecordNotFound)
	if isNotFound {
		return nil, res.Error
	}
	return user, nil
}

func (u *userRepository) ForgotPassword(info *models.User) (*models.User, error) {
	var user *models.User

	err := u.db.Where("ktp_id = ?", info.KTPID).Updates(&info).First(&user)
	if err.Error != nil {
		return nil, err.Error
	}

	return user, nil
}

func (u *userRepository) Location() ([]*models.Location, error) {
	var location []*models.Location

	res := u.db.Find(&location)

	if res.Error != nil {
		return nil, res.Error
	}

	return location, nil
}

func (u *userRepository) CreateAttendance(attendance *models.Attendance) (*models.Attendance, error) {
	res := u.db.Select("CheckInOut", "UserID", "LocationID").Create(&attendance)

	if res.Error != nil {
		return nil, res.Error
	}

	recordedAttendance := &models.Attendance{
		ID:         attendance.ID,
		UserID:     attendance.UserID,
		LocationID: attendance.LocationID,
	}

	return recordedAttendance, nil
}

func (u *userRepository) CreateCheckout(attendance *models.Attendance) (*models.Attendance, error) {
	res := u.db.Select("CheckInOut", "UserID", "LocationID").Create(&attendance)

	if res.Error != nil {
		return nil, res.Error
	}

	recordedAttendance := &models.Attendance{
		ID:         attendance.ID,
		UserID:     attendance.UserID,
		LocationID: attendance.LocationID,
	}

	return recordedAttendance, nil
}

func (u *userRepository) Logs(filterBy string) ([]*models.Attendance, error) {
	var attendance []*models.Attendance
	res := u.db

	if filterBy == "day" {
		res = u.db.Where("DATE(created_at) = CURDATE()").Find(&attendance)
	}

	if filterBy == "week" {
		res = u.db.Where("YEARWEEK(created_at, 1) = YEARWEEK(CURDATE(), 1)").Find(&attendance)
	}

	if filterBy == "month" {
		res = u.db.Where("MONTH(DATE(created_at)) = MONTH(CURDATE()) AND YEAR(DATE(created_at)) = YEAR(CURDATE())").Find(&attendance)
	}

	if filterBy == "Year" {
		res = u.db.Where("YEAR(DATE(created_at)) = YEAR(CURDATE())").Find(&attendance)
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return attendance, nil
}

func (u *userRepository) Profile(id int) (*models.User, error) {
	var user *models.User

	res := u.db.Where("id = ?", id).First(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}
