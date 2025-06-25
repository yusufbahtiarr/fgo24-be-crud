package routers

import (
	"fgo24-be-crud/controllers"

	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllUsers)
	r.GET("/:id", controllers.GetUserByID)
	r.POST("", controllers.CreateUser)
	r.PATCH("/:id", controllers.UpdateUser)
	r.DELETE("/:id", controllers.DeleteUser)

}
