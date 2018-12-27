package handler

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // ini untuk db
)

func getConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", os.Getenv("DBCONN"))
	if err != nil {
		panic(err)
	}

	return
}
