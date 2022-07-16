package controllers

import (
	"mime/multipart"
	"net/http"

	"github.com/VictorRibeiroLima/cloud-storage/models"
	"github.com/VictorRibeiroLima/cloud-storage/utils"
	"github.com/gin-gonic/gin"
)

type FileUploader interface {
	UploadFile(models.User, *multipart.FileHeader)
}

type FileLister interface {
	ListFiles(models.User) []struct {
		Id   uint
		File string
	}
}

type StorageController struct {
	FileUploader FileUploader
	FileLister   FileLister
}

func (c *StorageController) UploadFile(context *gin.Context) {
	file, _ := context.FormFile("file")
	user := utils.GetUser(context)
	c.FileUploader.UploadFile(user, file)
}

func (c *StorageController) ListFiles(context *gin.Context) {
	user := utils.GetUser(context)
	files := c.FileLister.ListFiles(user)
	context.JSON(http.StatusOK, files)
}
