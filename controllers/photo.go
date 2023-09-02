package controllers

import (
	"encoding/base64"
	"errors"
	"io"
	"mime/multipart"
	"path/filepath"

	"shopingCar_go/constants"
	"shopingCar_go/customerrors"

	"github.com/gin-gonic/gin"
)

func (c *Controller) UploadImage(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["avatar[]"]
	var base64Files []string
	for _, file := range files {
		if file.Size > constants.MaxFileSize {
			customerrors.ErrorResponse(string(constants.MaxFileError), ctx, errors.New("invalid file extension"))
			return
		}
		ext := filepath.Ext(file.Filename)
		var validExtension = false
		allowedExtensions := []string{".jpg", ".jpeg", "png"}
		for _, allowedExt := range allowedExtensions {
			if ext == allowedExt {
				validExtension = true
			}
		}
		if !validExtension {
			customerrors.ErrorResponse(string(constants.NotAllowedExtensionError), ctx, errors.New("extension not allowed"))
			return
		}
		base64File, err := convertFileToBase64(file)
		if err != nil {
			customerrors.ErrorResponse(string(constants.FileConversionError), ctx, errors.New("file conversion error"))
			return
		}
		base64Files = append(base64Files, base64File)
	}

}

func convertFileToBase64(file *multipart.FileHeader) (string, error) {
	openedFile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer openedFile.Close()

	buffer := make([]byte, file.Size)
	_, err = openedFile.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	base64String := base64.StdEncoding.EncodeToString(buffer)
	return base64String, nil
}
