package routers

import (
	"fgo24-be-crud/controllers"

	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.RouterGroup) {
	r.GET("", controllers.GetDataUser)
	r.GET("/:id", controllers.GetUserById)
	r.POST("", controllers.CreateUser)
	r.PATCH("/:id", controllers.UpdateUser)
	r.DELETE("/:id", controllers.DeleteUser)

}
