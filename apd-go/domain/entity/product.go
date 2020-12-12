package entity

import "time"

//Product entity
type Product struct {
	ID          uint64    `gorm:"primary_key;auto_increment"`
	Name        string    `gorm:"size:100;not null"`
	Description string    `gorm:"size:1000;not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
