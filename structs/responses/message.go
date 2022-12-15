package responses

import "time"

type Message struct {
	ID uint `json:"id"`
	ConversationID uint `json:"conversation_id"`
	MessageReplyID uint  `json:"message_reply_id"`
	Text      string  `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
