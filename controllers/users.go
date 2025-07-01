package controllers

import (
	"context"
	"encoding/json"
	"fgo24-be-crud/models"
	"fgo24-be-crud/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Description List All Users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} string "string"
// @Router /users [get]
func GetAllUsers(ctx *gin.Context) {
	err := utils.RedisClient.Ping(context.Background()).Err()
	noredis := false
	if err != nil {
		if strings.Contains(err.Error(), "refused") {
			noredis = true
		}
	}

	if !noredis {
		result := utils.RedisClient.Exists(context.Background(), ctx.Request.RequestURI)
		if result.Val() != 0 {
			users := models.User{}
			data := utils.RedisClient.Get(context.Background(), ctx.Request.RequestURI)
			str := data.Val()
			if err = json.Unmarshal([]byte(str), &users); err != nil {
				log.Println("Unmarshal error:", err)
			} else {
				ctx.JSON(http.StatusOK, utils.Response{
					Success: true,
					Message: "List all users (from Redis)",
					Results: users,
				})
			}
			return
		}
	}

	users, err := models.FindAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed show all users",
		})
		return
	}

	if !noredis {
		encoded, err := json.Marshal(users)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, utils.Response{
				Success: false,
				Message: "Failed to get user from database",
			})
			return
		}
		utils.RedisClient.Set(context.Background(), ctx.Request.RequestURI, string(encoded), 0)
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success show all users",
		Results: users,
	})
}

// @Descrition 	Detail User By ID
// @Tags 				users
// @Accept 			json
// @Produce 		json
// @Param 			id 		path int 	true 		"User ID"
// @Success 		200 	{string} 	string 	"string"
// @Router 			/users/{id} 		[get]
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

// @Descrition Delete User By ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "string"
// @Router /users/{id} [delete]
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

// @Description Create User
// @Tags users
// @Accept json
// @Produce json
// @Param users body models.CreateUserRequest true "User data"
// @Success 200 {object} models.CreateUserRequest "string"
// @Failure 400 {object} utils.Response
// @Router /users [post]
func CreateUser(ctx *gin.Context) {
	user := models.User{}
	ctx.ShouldBindJSON(&user)

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

// UpdateUser godoc
// @Summary Update user data
// @Description Update existing user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UpdateUserRequest true "User update data"
// @Success 200 {object} models.UpdateUserRequest
// @Failure 400 {object} utils.Response
// @Router /users/{id} [patch]
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
