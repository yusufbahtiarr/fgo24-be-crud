package controllers

import (
	"fgo24-be-crud/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UploadFile godoc
// @Summary Upload a file
// @Description Upload a file to server
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {object} utils.Response
// @Success 400 {object} utils.Response
// @Router /upload [post]
func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Failed to upload file",
		})
		return
	}

	if file.Size > 1*1024*1024 {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "File is too big",
		})
		return
	}

	newName := uuid.New().String() + filepath.Ext(file.Filename)
	err = ctx.SaveUploadedFile(file, "./upload/"+newName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to save file",
		})
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Upload Success.",
	})

}
