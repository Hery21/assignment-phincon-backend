package handlers

import (
	"GOLANG/dto"
	"net/http"

	"GOLANG/httperror"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LogIn(c *gin.Context) {
	println("masuk1")
	value, _ := c.Get("payload")
	logInReq := value.(*dto.LogInReq)

	result, err := h.authService.LogIn(logInReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}
	println("masuk2")

	c.JSON(http.StatusOK, result)
}
