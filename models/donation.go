package models

import "time"


type Donation struct {
	ID        uint      `gorm:"primaryKey"`
	Amount    int       `gorm:"not null"`
	Status    string    `gorm:"type:varchar(20);not null"`
	Authority string    `gorm:"type:varchar(50);unique"`
	RefID     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
