package models

// User model
type User struct {
	ID       string `gorm:"primaryKey"`
	Email    string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Role     string `gorm:"not null"`
}
