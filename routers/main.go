package routers

import (
	"fgo24-be-crud/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// docs "github.com/yusufbahtiarr/fgo24-be-crud/docs"
)

func CombineRouter(r *gin.Engine) {
	userRoute(r.Group("/users"))
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/docs", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/docs/index.html")
	})
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
