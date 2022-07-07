package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database interface {
	//OpenDB()
	CloseDB()
	AutoMigrate(values ...interface{})
}

type database struct {
	conn *gorm.DB
}

func NewRepository() Database {
	con, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to open database")
	}

	return &database{
		conn: con,
	}
}

// func (db *database) OpenDB() {
// 	con, err := gorm.Open("sqlite3", "test.db")
// 	if err != nil {
// 		panic("Failed to open database")
// 	}

// 	db.conn = con
// }

func (db *database) CloseDB() {
	err := db.conn.Close()
	if err != nil {
		panic("Failed to close database")
	}

	db.conn = nil
}

func (db *database) AutoMigrate(values ...interface{}) {
	db.conn.AutoMigrate(values)
}
