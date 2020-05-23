package main

import (
	conn "CRUD/server/conn"
	middleware "CRUD/server/middleware"
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
	s.ListenAndServe()
}

//Function to check Status of the server
func checkStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is running successfully !!!!!")
	}
}
