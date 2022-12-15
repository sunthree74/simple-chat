package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ConversationID uint `json:"conversation_id"`
	MessageReplyID uint  `json:"message_reply_id"`
	Text      string  `json:"text"`
}

func (Message) TableName() string {
	return "messages"
}
