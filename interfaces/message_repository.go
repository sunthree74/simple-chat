package interfaces

import (
	"context"
	"github.com/sunthree74/simple-chat/model"
	"github.com/sunthree74/simple-chat/structs/responses"
)

type MessageRepository interface {
	GetByConversationID(ctx context.Context, conversationID uint) ([]responses.Message, error)
	Create(ctx context.Context, message model.Message) error
}
