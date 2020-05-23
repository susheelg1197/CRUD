package routes

import (
	handlers "CRUD/server/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(o *gin.RouterGroup) {
	o.GET("blood-bank/show-details", handlers.ShowUserDetails())
	o.POST("blood-bank/add-user-details", handlers.AddUserDetails())
	o.POST("blood-bank/update-user-details", handlers.UpdateUserDetails())
}
