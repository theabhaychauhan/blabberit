package message

import (
	"time"
)

type Message struct {
	ID        uint `gorm:"primaryKey"`
	From      string
	To        string
	Content   string
	Timestamp time.Time
}
