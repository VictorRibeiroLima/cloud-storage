package models

type Storage struct {
	Model
	FileName     string
	FileExtesion string
	File         []byte
	UserId       uint
	User         User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
