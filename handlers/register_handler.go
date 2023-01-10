package handlers

import (
	"GOLANG/dto"
	"GOLANG/httperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) Register(c *gin.Context) {
	value, _ := c.Get("payload")
	registerReq := value.(*dto.RegisterReq)

	result, err := h.registerService.Register(registerReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
}
