package middlewares

import (
	"GOLANG/httperror"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ModelCreator func() any

func RequestValidator(creator ModelCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		model := creator()
		c.Set("payload", model)

		if err := c.ShouldBindJSON(&model); err != nil {
			fmt.Println(err.Error())
			badRequestError := httperror.BadRequestError(err.Error(), "BAD_REQUEST")
			c.AbortWithStatusJSON(badRequestError.StatusCode, badRequestError)
		}

		c.Set("payload", model)
	}
}
