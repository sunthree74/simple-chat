package interfaces

import (
	"context"
	"github.com/sunthree74/simple-chat/model"
)

type ConversationUsecase interface {
	Create(ctx context.Context, conversation model.Conversation) (err error, cnv model.Conversation)
	GetByUserID(ctx context.Context, userID uint) ([]model.Conversation, error)
	GetByReceiverID(ctx context.Context, receiverID uint) (model.Conversation, error)
	EmptyUnreadCount(ctx context.Context, id uint) error
}
