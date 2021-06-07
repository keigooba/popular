package models

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	dbIp := os.Getenv("DB_IP")
	var PROTOCOL string
	if dbIp != "" {
		PROTOCOL = "tcp(" + dbIp + ":3306)"
	} else {
		PROTOCOL = "tcp(127.0.0.1:3306)"
	}
	DBNAME := "gomix_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	_, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		log.Fatal(err)
	}
}
