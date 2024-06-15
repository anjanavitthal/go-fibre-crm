package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDatabase() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		fmt.Println("Failed to connect to database")
	}

	fmt.Println("Database connection is open")

}

var (
	DBConn *gorm.DB
)
