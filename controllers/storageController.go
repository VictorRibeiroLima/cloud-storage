package controllers

import (
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/VictorRibeiroLima/cloud-storage/models"
	responsebuilder "github.com/VictorRibeiroLima/cloud-storage/response-builder"
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

type FileFinder interface {
	FindFile(models.User, uint) (models.Storage, error)
}

type StorageController struct {
	FileUploader FileUploader
	FileLister   FileLister
	FileFinder   FileFinder
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

func (c *StorageController) Download(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))

	user := utils.GetUser(context)
	file, err := c.FileFinder.FindFile(user, uint(id))
	if err != nil {
		responsebuilder.NotFound(context, "Storage")
		return
	}
	byteFile := file.File
	fileName := file.FileName + "." + file.FileExtesion
	context.Header("Content-Disposition", "attachment; filename="+fileName)
	context.Data(http.StatusOK, "application/octet-stream", byteFile)
}
