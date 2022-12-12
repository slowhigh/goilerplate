package models

// Memo model
type Memo struct {
	ID      string `gorm:"primaryKey"`
	Content string `gorm:"not null"`
	UserID  string `gorm:"not null"`
}
