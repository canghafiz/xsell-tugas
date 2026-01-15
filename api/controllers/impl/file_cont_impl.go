package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileContImpl struct {
	FileServ services.FileServ
}

func NewFileContImpl(fileServ services.FileServ) *FileContImpl {
	return &FileContImpl{FileServ: fileServ}
}

func (cont *FileContImpl) UploadFiles(context *gin.Context) {
	form, err := context.MultipartForm()
	if err != nil {
		exceptions.ErrorHandler(context, err)
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		exceptions.ErrorHandler(context, "no files uploaded")
		return
	}

	uploadedFiles, errServ := cont.FileServ.UploadFiles(files)
	if errServ != nil {
		exceptions.ErrorHandler(context, errServ)
		return
	}

	api := helpers.ApiResponse{
		Success: true,
		Code:    http.StatusOK,
		Data:    uploadedFiles,
	}

	if err := helpers.WriteToResponseBody(context, api.Code, api); err != nil {
		exceptions.ErrorHandler(context, err)
		return
	}
}

func (cont *FileContImpl) DeleteFiles(context *gin.Context) {
	url := context.Query("url")

	// Call Service
	err := cont.FileServ.DeleteFile(url)
	if err != nil {
		exceptions.ErrorHandler(context, err)
		return
	}

	api := helpers.ApiResponse{
		Success: true,
		Code:    http.StatusOK,
		Data:    nil,
	}

	if err := helpers.WriteToResponseBody(context, api.Code, api); err != nil {
		exceptions.ErrorHandler(context, err)
		return
	}
}
