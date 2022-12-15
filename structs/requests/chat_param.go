package requests

type ChatParam struct {
	ReceiverID uint `json:"receiver_id" binding:"required"`
	Text string `json:"text" binding:"required"`
	ConversationID uint `json:"conversation_id"`
	MessageReplyID uint `json:"message_reply_id"`
}
