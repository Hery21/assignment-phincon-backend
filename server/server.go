package server

import (
	"GOLANG/config"
	"GOLANG/db"
	"GOLANG/repositories"
	"GOLANG/services"
	"fmt"
)

func Init() {
	registerRepository := repositories.NewRegisterRepository(&repositories.RRConfig{DB: db.Get()})
	registerService := services.NewRegisterService(&services.RSConfig{RegisterRepository: registerRepository})

	userRepository := repositories.NewUserRepository(&repositories.URConfig{DB: db.Get()})
	authService := services.NewAuthService(
		&services.AuthSConfig{
			UserRepository: userRepository,
			AppConfig:      config.Config,
		})

	router := NewRouter(&RouterConfig{
		RegisterService: registerService,
		AuthService:     authService,
	})

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("server Error: ", err)
	}
}
