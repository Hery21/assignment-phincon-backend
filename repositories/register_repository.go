package repositories

import (
	"GOLANG/models"

	"gorm.io/gorm"
)

type RegisterRepository interface {
	Register(user *models.User) (*models.User, error)
}

type registerRepository struct {
	db *gorm.DB
}

type RRConfig struct {
	DB *gorm.DB
}

func NewRegisterRepository(c *RRConfig) *registerRepository {
	return &registerRepository{db: c.DB}
}

func (r *registerRepository) Register(user *models.User) (*models.User, error) {
	res := r.db.Select("Username", "FullName", "KTPID", "Password").Create(&user)

	if err := res.Error; err != nil {
		return &models.User{}, err
	}

	registeredUser := &models.User{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		KTPID:    user.KTPID,
	}

	return registeredUser, res.Error
}
