package repository

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type MySQL interface {
	CloseDB()
	AutoMigrate(...interface{})

	Create(interface{})
	Save(interface{})
	Delete(interface{})
	FindAll(interface{})
}

type _mysql struct {
	conn *gorm.DB
}


func NewMySQL() MySQL {
	dsn := "root:1234@tcp(127.0.0.1:3306)/yt_go_auth?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}

	return &_mysql{
		conn: db,
	}
}

func (db *_mysql) AutoMigrate(values ...interface{}) {
	db.conn.AutoMigrate(values...)
}

func (db *_mysql) CloseDB() {
	// err := db.conn.
	// if err != nil {
	// 	panic("Failed to close _mysql")
	// }

	// db.conn = nil
}

func (db *_mysql) Create(value interface{}) {
	db.conn.Create(value)
}

func (db *_mysql) Save(value interface{}) {
	db.conn.Save(value)
}

func (db *_mysql) Delete(value interface{}) {
	db.conn.Delete(value)
}

func (db *_mysql) FindAll(value interface{}) {
	db.conn.Set("gorm:auto_preload", true).Find(value)
}