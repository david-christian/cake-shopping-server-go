package controllers

import (
	"errors"
	"net/http"
	"path/filepath"

	"shopingCar_go/constants"
	"shopingCar_go/customerrors"

	"github.com/gin-gonic/gin"
)

func (c *Controller) UploadImage(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["avatar"]
	for _, file := range files {
		if file.Size > constants.MaxFileSize {
			customerrors.ErrorResponse(string(constants.MaxFileError), ctx, errors.New("invalid file extension"))
			return
		}
		ext := filepath.Ext(file.Filename)
		var validExtension = false
		allowedExtensions := []string{".jpg", ".jpeg", ".png"}
		for _, allowedExt := range allowedExtensions {
			if ext == allowedExt {
				validExtension = true
			}
		}
		if !validExtension {
			customerrors.ErrorResponse(string(constants.NotAllowedExtensionError), ctx, errors.New("extension not allowed"))
			return
		}
	}
	_, err := c.service.UploadImage(files)
	if err != nil {
		customerrors.ErrorResponse(string(constants.TypeError), ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
