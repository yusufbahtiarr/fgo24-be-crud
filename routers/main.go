package routers

import "github.com/gin-gonic/gin"

func CombineRouter(r *gin.Engine) {
	userRoute(r.Group("/users"))
}
