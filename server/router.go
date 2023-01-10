package server

import (
	"GOLANG/dto"
	"GOLANG/handlers"
	"GOLANG/middlewares"
	"GOLANG/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	RegisterService services.RegisterService
	AuthService     services.AuthService
}

func NewRouter(c *RouterConfig) *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{
		"Access-Control-Allow-Headers",
		"Authorization",
		"Origin",
		"Accept",
		"X-Requested-With",
		"Content-Type",
		"Access-Control-Request-Method",
		"Access-Control-Request-Headers",
	}
	router.Use(cors.New(config))

	h := handlers.New(&handlers.HandlerConfig{
		RegisterService: c.RegisterService,
		AuthService:     c.AuthService,
	})

	router.POST("/register", middlewares.RequestValidator(func() any {
		return &dto.RegisterReq{}
	}), h.Register, middlewares.ErrorHandler)

	router.POST("/login", middlewares.RequestValidator(func() any {
		return &dto.LogInReq{}
	}), h.LogIn, middlewares.ErrorHandler)

	router.PATCH("/forgot-password", middlewares.RequestValidator(func() any {
		return &dto.ForgotPasswordReq{}
	}), h.ForgotPassword, middlewares.ErrorHandler)

	router.GET("/attendance", middlewares.AuthorizeJWT, h.Locations, middlewares.ErrorHandler)

	router.POST("/check-in/:location_id", middlewares.AuthorizeJWT, h.CreateAttendance, middlewares.ErrorHandler)

	router.POST("/check-out/:location_id", middlewares.AuthorizeJWT, h.CreateCheckout, middlewares.ErrorHandler)

	router.GET("/log", middlewares.AuthorizeJWT, h.Logs, middlewares.ErrorHandler)

	router.GET("/profile", middlewares.AuthorizeJWT, h.Profile, middlewares.ErrorHandler)

	return router
}
