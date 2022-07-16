package controllers

import (
	"mime/multipart"

	"github.com/VictorRibeiroLima/cloud-storage/models"
	"github.com/VictorRibeiroLima/cloud-storage/utils"
	"github.com/gin-gonic/gin"
)

type FileUploader interface {
	UploadFile(models.User, *multipart.FileHeader)
}

type StorageController struct {
	FileUploader FileUploader
}

func (c *StorageController) UploadFile(context *gin.Context) {
	file, _ := context.FormFile("file")
	user := utils.GetUser(context)
	c.FileUploader.UploadFile(user, file)
}
