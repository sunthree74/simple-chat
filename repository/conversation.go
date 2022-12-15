package repository

import (
	"context"
	"fmt"
	"github.com/sunthree74/simple-chat/interfaces"
	"github.com/sunthree74/simple-chat/model"
	"github.com/sunthree74/simple-chat/structs/responses"
	"gorm.io/gorm"
)

var _ interfaces.ConversationRepository = (*conversationRepository)(nil)

type conversationRepository struct {
	db *gorm.DB
}

func (c *conversationRepository) EmptyUnreadCount(ctx context.Context, id uint) error {
	var conversation model.Conversation
	err := c.db.WithContext(ctx).
		Model(conversation).
		Where("id = ?", id).
		Update("unread_count", 0).
		Error
	if err != nil {
		return fmt.Errorf("empty unread count by id query error: %w", err)
	}

	return nil
}

func (c *conversationRepository) IncrementUnreadCount(ctx context.Context, id uint, lastMessageID uint) error {
	err := c.db.WithContext(ctx).
		Exec("UPDATE conversations SET unread_count = unread_count + 1, last_message_id = ? WHERE id = ?", lastMessageID, id).
		Error
	if err != nil {
		return fmt.Errorf("increment unread_count by id query error: %w", err)
	}

	return nil
}

func (c *conversationRepository) GetByUserID(ctx context.Context, userID uint) ([]responses.Conversation, error) {
	var conversations []responses.Conversation
	err := c.db.WithContext(ctx).
		Model(model.Conversation{}).
		Select("users.name, conversations.*, messages.*").
		Joins("left join users on conversations.receiver_id = users.id").
		Joins("left join messages on conversations.last_message_id = messages.id").
		Where("user_id = ?", userID).
		Scan(&conversations).
		Error
	if err != nil {
		return []responses.Conversation{}, fmt.Errorf("get conversation by user id query error: %w", err)
	}

	if len(conversations) < 1 {
		return []responses.Conversation{}, gorm.ErrRecordNotFound
	}

	return conversations, nil
}

func (c *conversationRepository) GetByReceiverID(ctx context.Context, receiverID uint) (model.Conversation, error) {
	var conversation model.Conversation
	err := c.db.WithContext(ctx).
		Where("receiver_id = ?", receiverID).
		Find(&conversation).
		Error
	if err != nil {
		return model.Conversation{}, fmt.Errorf("get conversation by user id query error: %w", err)
	}

	return conversation, nil
}

func (c *conversationRepository) Create(ctx context.Context, conversation model.Conversation) (error, model.Conversation) {
	if err := c.db.WithContext(ctx).Save(&conversation).Error; err != nil {
		return fmt.Errorf("save query error: %w", err), model.Conversation{}
	}

	return nil, conversation
}

func InitializeConversation(db *gorm.DB) *conversationRepository {
	return &conversationRepository{db: db}
}
