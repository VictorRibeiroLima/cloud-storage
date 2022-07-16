package services

import (
	"bytes"
	"io"
	"mime/multipart"
	"strings"

	"github.com/VictorRibeiroLima/cloud-storage/models"
	"gorm.io/gorm"
)

type StorageService struct {
	Db *gorm.DB
}

func (s *StorageService) UploadFile(user models.User, fileHeader *multipart.FileHeader) {
	blob := []byte{}
	buf := bytes.NewBuffer(blob)
	filePath := fileHeader.Filename
	fileInfos := strings.Split(filePath, ".")
	file, openErr := fileHeader.Open()
	if openErr != nil {
		panic("Error opening file")
	}
	if _, err := io.Copy(buf, file); err != nil {
		panic("Error buffering file")
	}
	storage := models.Storage{
		File:         buf.Bytes(),
		FileName:     fileInfos[0],
		FileExtesion: fileInfos[1],
		User:         user,
	}
	s.Db.Create(&storage)
}

func (s *StorageService) ListFiles(user models.User) []struct {
	Id   uint
	File string
} {
	var files []struct {
		Id   uint
		File string
	}
	rows, _ := s.Db.Model(&models.Storage{}).Where("user_id = ?", user.ID).Rows()
	defer rows.Close()

	for rows.Next() {
		var file models.Storage

		s.Db.ScanRows(rows, &file)
		files = append(files, struct {
			Id   uint
			File string
		}{
			Id:   file.ID,
			File: file.FileName + "." + file.FileExtesion,
		})
	}
	return files
}
