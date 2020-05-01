package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowUserDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Successfull")
	}
}
