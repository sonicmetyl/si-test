package adapters

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"go-sql/app/helpers"
)

var (
	DbReadConn  *gorm.DB
	DbWriteConn *gorm.DB
)

type DbAdapter struct {
	logMode                  bool
	writerConn, readerConn   string
}

func (dba *DbAdapter) Init() {
	dba.writerConn = helpers.GetConfigString("DB.WRITE")
	dba.readerConn = helpers.GetConfigString("DB.READ")
	dba.logMode = helpers.GetConfigBool("DEBUG")

	dba.GetWriteDB()
	dba.GetReadDB()
}

func (dba *DbAdapter) GetWriteDB() *gorm.DB {
	if DbWriteConn == nil {
		DbWriteConn = dba.createDBConnection(dba.writerConn)
	}

	return DbWriteConn
}

func (dba *DbAdapter) GetReadDB() *gorm.DB {
	if DbReadConn == nil {
		DbReadConn = dba.createDBConnection(dba.readerConn)
	}

	return DbReadConn
}

func (dba *DbAdapter) createDBConnection(descriptor string) *gorm.DB {
	db, err := gorm.Open("mysql", descriptor)
	if err != nil {
		log.Println("Error connecting to DB: ", err)
		os.Exit(1)
	}

	db.LogMode(dba.logMode)
	return db
}
