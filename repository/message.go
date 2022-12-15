package repository

import (
	"context"
	"fmt"
	"github.com/sunthree74/simple-chat/interfaces"
	"github.com/sunthree74/simple-chat/model"
	"github.com/sunthree74/simple-chat/structs/responses"
	"gorm.io/gorm"
)

var _ interfaces.MessageRepository = (*messageRepository)(nil)

type messageRepository struct {
	db *gorm.DB
}

func (m *messageRepository) GetByConversationID(ctx context.Context, conversationID uint) ([]responses.Message, error) {
	var messages []responses.Message
	err := m.db.WithContext(ctx).
		Model(model.Message{}).
		Where("conversation_id = ?", conversationID).
		Scan(&messages).
		Error
	if err != nil {
		return []responses.Message{}, fmt.Errorf("get messages by conversation id query error: %w", err)
	}

	if len(messages) < 1 {
		return []responses.Message{}, gorm.ErrRecordNotFound
	}

	return messages, nil
}

func (m *messageRepository) Create(ctx context.Context, message model.Message) (model.Message, error) {
	if err := m.db.WithContext(ctx).Save(&message).Error; err != nil {
		return model.Message{}, nil
	}

	return message, nil
}

func InitializeMessage(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}
