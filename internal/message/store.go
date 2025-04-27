package message

import "gorm.io/gorm"

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
}

func SaveMessage(m *Message) error {
	return DB.Create(m).Error
}

func GetMessagesForUser(publicKey string) ([]Message, error) {
	var messages []Message
	err := DB.Where(`"to" = ?`, publicKey).
		Order("timestamp ASC").
		Find(&messages).Error

	return messages, err
}
