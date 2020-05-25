package handlers

import (
	"CRUD/server/models"
	"CRUD/server/services"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

func UpdateDocumentDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := models.UserDetails{}
		c.Bind(&requestBody)
		services.UpdateDocuments(requestBody)
		c.JSON(http.StatusOK, "Successfull")
	}
}

func DeleteUserDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := models.UserDetails{}
		c.Bind(&requestBody)
		services.DeleteUsers(requestBody)
		c.JSON(http.StatusOK, "Successfull")
	}
}

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
			return
		}
		filename := header.Filename
		out, err := os.Create("public/" + filename)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
			return
		}
		filepath := "http://localhost:8787/file/" + filename
		c.JSON(http.StatusOK, gin.H{"filepath": filepath})
	}
}
