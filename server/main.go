package main

import (
	conn "CRUD/server/conn"
	middleware "CRUD/server/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	conn.InitConn()
	startServer()
}

//Function to Start Server
func startServer() {
	router := gin.Default()
	md := cors.DefaultConfig()
	md.AllowAllOrigins = true
	md.AllowHeaders = []string{"*"}
	md.AllowMethods = []string{"*"}
	router.Use(cors.New(md))
	middleware.InitMiddleware(router)
	router.GET("/", checkStatus())
	s := &http.Server{
		Addr:    ":8787",
		Handler: router,
	}
	router.POST("/one/:formNo", func(c *gin.Context) {
		// single file
		formNo := c.Param("formNo")
		file, err := c.FormFile("file")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(file.Filename)
		err = c.SaveUploadedFile(file, "file/profilePictures/"+formNo+".png")
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Static("/file/profilePictures", "./file/profilePictures")
	s.ListenAndServe()
}

//Function to check Status of the server
func checkStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is running successfully !!!!!")
	}
}
