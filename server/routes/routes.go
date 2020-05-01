package routes

import (
	handlers "Crud-App/server/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(o *gin.RouterGroup) {
	o.GET("blood_bank/show_details", handlers.ShowUserDetails())

}
