package gomud

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // Needed by gorm
	"log"
)

// GetDb returns a gorm connection
func GetDb() gorm.DB {
	db, err := gorm.Open("postgres", "user=dan dbname=gomud sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
