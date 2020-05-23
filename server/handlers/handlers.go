package handlers

import (
	"CRUD/server/models"
	"CRUD/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowUserDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		list := services.ShowUsersList()
		c.JSON(http.StatusOK, list)
	}
}

func AddUserDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := models.UserDetails{}
		c.Bind(&requestBody)
		services.AddUsers(requestBody)
		c.JSON(http.StatusOK, "Successfull")
	}
}

func UpdateUserDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := models.UserDetails{}
		c.Bind(&requestBody)
		services.UpdateUsers(requestBody)
		c.JSON(http.StatusOK, "Successfull")
	}
}
