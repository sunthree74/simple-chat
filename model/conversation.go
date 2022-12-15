package model

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	UserID uint  `json:"user_id"`
	ReceiverID      uint  `json:"receiver_id"`
	UnreadCount int `json:"unread_count"`
}

func (Conversation) TableName() string {
	return "conversations"
}
