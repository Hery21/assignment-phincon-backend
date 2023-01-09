package middlewares

import (
	"GOLANG/httperror"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	if len(c.Errors) == 0 {
		return
	}

	firstError := c.Errors[0].Err
	fmt.Println("ErrorHandler: ", firstError)

	appError, isAppError := firstError.(httperror.AppError)
	if isAppError {
		c.JSON(appError.StatusCode, appError)
		return
	}

	serverErr := httperror.InternalServerError(firstError.Error())
	c.JSON(serverErr.StatusCode, serverErr)
}
