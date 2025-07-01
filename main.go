package main

import (
	"fgo24-be-crud/routers"
	"fgo24-be-crud/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @Title 			CRUD Users
// @Version 		1.0
// @Description	This Simple CRUD server
// @BasePath /

func main() {
	r := gin.Default()
	godotenv.Load()
	routers.CombineRouter(r)
	utils.Redis()

	r.Run(":" + os.Getenv("APP_PORT"))
}
