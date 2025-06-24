package main

import (
	"fgo24-be-crud/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	godotenv.Load()
	routers.CombineRouter(r)

	r.Run()
}
