package mysql

import (
	"github.com/jinzhu/gorm"
	"go-sql/app/repositories"
)

type Storage struct {
	readConn *gorm.DB
	writeConn *gorm.DB
}

func NewStorage(readConn, writeConn *gorm.DB) repositories.StorageInterface {
	return &Storage{
		readConn:readConn,
		writeConn:writeConn,
	}
}