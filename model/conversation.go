package model

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	UserID        uint `json:"user_id"`
	ReceiverID    uint `json:"receiver_id"`
	UnreadCount   int  `json:"unread_count"`
	LastMessageID uint `json:"last_message_id"`
}

func (Conversation) TableName() string {
	return "conversations"
}
