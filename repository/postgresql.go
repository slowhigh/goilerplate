package repository

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type PostgreSQL interface {
	CloseDB()
	AutoMigrate(...interface{})

	Create(interface{})
	Save(interface{})
	Delete(interface{})
	FindAll(interface{})
}

type postgreSQL struct {
	conn *gorm.DB
}


func NewPostgreSQL() PostgreSQL {
	dsn := "host=127.0.0.1 user=postgres password=1234 dbname=yt_go_auth port=5433 sslmode=disable TimeZone=Asia/Shanghai"
  	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}

	return &postgreSQL{
		conn: db,
	}
}

func (db *postgreSQL) AutoMigrate(values ...interface{}) {
	db.conn.AutoMigrate(values...)
}

func (db *postgreSQL) CloseDB() {
	// err := db.conn.
	// if err != nil {
	// 	panic("Failed to close postgreSQL")
	// }

	// db.conn = nil
}

func (db *postgreSQL) Create(value interface{}) {
	db.conn.Create(value)
}

func (db *postgreSQL) Save(value interface{}) {
	db.conn.Save(value)
}

func (db *postgreSQL) Delete(value interface{}) {
	db.conn.Delete(value)
}

func (db *postgreSQL) FindAll(value interface{}) {
	db.conn.Set("gorm:auto_preload", true).Find(value)
}