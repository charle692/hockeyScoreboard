package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // imported for use by gorm
)

// Conn - Connection to the database
var Conn *gorm.DB

func init() {
	var err error
	Conn, err := gorm.Open("sqlite3", "./hockeyScoreboard.db")

	if err != nil {
		Close()
		log.Fatal("Failed to open database: ", err)
	}

	Conn.LogMode(true)
}

// Close - Closes the database connection
func Close() {
	if Conn != nil {
		Conn.Close()
	}
}
