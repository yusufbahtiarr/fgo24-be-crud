package controllers

import (
	"fgo24-be-crud/models"
	"fgo24-be-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDataUser(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")
	users, _ := models.AllUser(search)
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "List User",
		Results: users,
	})

}

func GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	users, _ := models.UserById(id)
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "List User",
		Results: users,
	})
}

func CreateUser(ctx *gin.Context) {
}

func UpdateUser(ctx *gin.Context) {
}

func DeleteUser(ctx *gin.Context) {
}
