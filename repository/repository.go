package repository

import (
	"github.com/jinzhu/gorm"
)

type Database interface {
	//OpenDB()
	
	CloseDB()
	AutoMigrate(...interface{})

	Create(interface{})
	Save(interface{})
	Delete(interface{})
	FindAll(interface{})
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

func (db *database) AutoMigrate(values ...interface{}) {
	db.conn.AutoMigrate(values...)
}

func (db *database) CloseDB() {
	err := db.conn.Close()
	if err != nil {
		panic("Failed to close database")
	}

	db.conn = nil
}

func (db *database) Create(value interface{}) {
	db.conn.Create(value)
}

func (db *database) Save(value interface{}) {
	db.conn.Save(value)
}

func (db *database) Delete(value interface{}) {
	db.conn.Delete(value)
}

func (db *database) FindAll(value interface{}) {
	db.conn.Set("gorm:auto_preload", true).Find(value)
}