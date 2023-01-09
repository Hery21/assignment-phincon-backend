package handlers

import (
	"GOLANG/dto"
	"GOLANG/httperror"
	"GOLANG/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ForgotPassword(c *gin.Context) {
	payload, _ := c.Get("payload")
	profileReq := payload.(*dto.ForgotPasswordReq)

	user, err := h.authService.ForgotPassword(profileReq)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) Locations(c *gin.Context) {
	locations, err := h.authService.Locations()
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, locations)
}

func (h *Handler) CreateAttendance(c *gin.Context) {
	user, _ := c.Get("user")
	ur := user.(models.User)
	location_id, _ := strconv.Atoi(c.Param("location_id"))

	result, err := h.authService.CreateAttendance(ur.ID, location_id)
	if err != nil {
		_ = c.Error(httperror.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) CreateCheckout(c *gin.Context) {
	user, _ := c.Get("user")
	ur := user.(models.User)
	location_id, _ := strconv.Atoi(c.Param("location_id"))

	result, err := h.authService.CreateCheckout(ur.ID, location_id)
	if err != nil {
		_ = c.Error(httperror.InternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) Logs(c *gin.Context) {
	filterBy := c.Query("filter-by")

	logs, err := h.authService.Logs(filterBy)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, logs)
}

func (h *Handler) Profile(c *gin.Context) {
	user, _ := c.Get("user")
	ur := user.(models.User)

	user, err := h.authService.Profile(ur.ID)
	if err != nil {
		_ = c.Error(httperror.InternalServerError("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, user)
}
