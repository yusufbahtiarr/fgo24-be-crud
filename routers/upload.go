package routers

import (
	"fgo24-be-crud/controllers"

	"github.com/gin-gonic/gin"
)

func uploadRouter(r *gin.RouterGroup) {
	r.POST("", controllers.UploadFile)
}
