package interfaces

import (
	"context"
	"github.com/sunthree74/simple-chat/model"
	"github.com/sunthree74/simple-chat/structs/responses"
)

type ConversationRepository interface {
	GetByUserID(ctx context.Context, userID uint) ([]responses.Conversation, error)
	GetByReceiverID(ctx context.Context, receiverID uint) (model.Conversation, error)
	Create(ctx context.Context, conversation model.Conversation) (error, model.Conversation)
	IncrementUnreadCount(ctx context.Context, id uint, lastMessageID uint) error
	EmptyUnreadCount(ctx context.Context, id uint) error
}
