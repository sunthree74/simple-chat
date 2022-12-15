package usecase

import (
	"context"
	"github.com/sunthree74/simple-chat/interfaces"
	"github.com/sunthree74/simple-chat/model"
	"github.com/sunthree74/simple-chat/structs/responses"
)

var _ interfaces.MessageUsecase = (*messageUsecase)(nil)

type messageUsecase struct {
	messageRepo      interfaces.MessageRepository
	conversationRepo interfaces.ConversationRepository
}

func (m *messageUsecase) GetByConversationID(ctx context.Context, conversationID uint) ([]responses.Message, error) {
	var messages []responses.Message
	messages, err := m.messageRepo.GetByConversationID(ctx, conversationID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (m *messageUsecase) Create(ctx context.Context, message model.Message) error {
	msg, err := m.messageRepo.Create(ctx, message)
	if err != nil {
		return nil
	}

	err = m.conversationRepo.IncrementUnreadCount(ctx, message.ConversationID, msg.ID)
	if err != nil {
		return nil
	}

	return nil
}

func InitializeMessage(messageRepo interfaces.MessageRepository, conversationRepo interfaces.ConversationRepository) *messageUsecase {
	return &messageUsecase{messageRepo: messageRepo, conversationRepo: conversationRepo}
}
