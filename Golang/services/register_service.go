package services

import (
	"GOLANG/dto"
	"GOLANG/models"
	"GOLANG/repositories"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService interface {
	Register(req *dto.RegisterReq) (*dto.RegisterRes, error)
}

type registerService struct {
	registerRepository repositories.RegisterRepository
}

type RSConfig struct {
	RegisterRepository repositories.RegisterRepository
}

func NewRegisterService(r *RSConfig) RegisterService {
	return &registerService{
		registerRepository: r.RegisterRepository,
	}
}

func (r *registerService) Register(req *dto.RegisterReq) (*dto.RegisterRes, error) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	req.Password = string(bytes)
	registeringUser := &models.User{
		Username: req.Username,
		FullName: req.FullName,
		Password: req.Password,
		KTPID:    req.KTPID,
	}

	registeredUser, err := r.registerRepository.Register(registeringUser)

	if err != nil {
		return new(dto.RegisterRes), err
	}

	return new(dto.RegisterRes).FromRegister(registeredUser), err
}
