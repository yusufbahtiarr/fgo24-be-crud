package controllers

import (
	"fgo24-be-crud/models"
	"fgo24-be-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed show all users",
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success show all users",
		Results: users,
	})
}

func GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := models.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed show user by id",
		})
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success show user by id",
		Results: user,
	})
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := models.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed Delete User by Id",
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success Deleted User by Id",
	})

}

func CreateUser(ctx *gin.Context) {
	user := models.User{}
	ctx.ShouldBind(&user)

	if user.Username == "" || user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Username or Email or Password cannot empty",
		})
		return
	}

	err := models.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed Create User",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success Create User",
	})
}

func UpdateUser(ctx *gin.Context) {
}
