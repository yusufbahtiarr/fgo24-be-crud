package controllers

import (
	"fgo24-be-crud/models"
	"fgo24-be-crud/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Description List All Users
// @Tags json
// @Accept json
// @Produce json
// @Success 200 {string} string "string"
// @Router /users [get]
func GetAllUsers(ctx *gin.Context) {
	users, err := models.FindAllUsers()
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

// @Descrition Detail User By ID
// @Tags json
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "string"
// @Router /users/{id} [get]
func GetUserByID(ctx *gin.Context) {
	idx := ctx.Param("id")
	id, err := strconv.Atoi(idx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	user, err := models.FindUserByID(id)
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
	idx := ctx.Param("id")
	id, err := strconv.Atoi(idx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	newData := models.User{}
	err = ctx.ShouldBind(&newData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	err = models.UpdateUser(id, newData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed Update User",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success Update User",
	})
}
