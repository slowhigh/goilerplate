package models

// User model
type User struct {
	ID       int64  `gorm:"primaryKey"`
	Email    string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
}
