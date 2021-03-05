package datastore

import (
	"log"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/yasuoza/clean-architecture-go/server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	mySqlConfig := &mysqlDriver.Config{
		User:                 config.C.Database.User,
		Passwd:               config.C.Database.Password,
		Net:                  config.C.Database.Net,
		Addr:                 config.C.Database.Addr,
		DBName:               config.C.Database.DBName,
		AllowNativePasswords: config.C.Database.AllowNativePasswords,
		ParseTime:            config.C.Database.ParseTime,
		Params: map[string]string{
			"charset": config.C.Database.Params.CharSet,
		},
	}

	db, err := gorm.Open(mysql.Open(mySqlConfig.FormatDSN()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
