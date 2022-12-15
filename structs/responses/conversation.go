package responses

import "time"

type Conversation struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	UserID         uint      `json:"user_id"`
	ReceiverID     uint      `json:"receiver_id"`
	UnreadCount    int       `json:"unread_count"`
	LastMessageID  uint      `json:"last_message_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
	ConversationID uint      `json:"conversation_id"`
	MessageReplyID uint      `json:"message_reply_id"`
	Text           string    `json:"text"`
}
