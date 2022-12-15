package usecase

import (
	"context"
	"github.com/sunthree74/simple-chat/interfaces"
	"github.com/sunthree74/simple-chat/model"
	"github.com/sunthree74/simple-chat/structs/responses"
)

var _ interfaces.ConversationUsecase = (*conversationUsecase)(nil)

type conversationUsecase struct {
	conversationRepo interfaces.ConversationRepository
}

func (c *conversationUsecase) EmptyUnreadCount(ctx context.Context, id uint) error {
	err := c.conversationRepo.EmptyUnreadCount(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *conversationUsecase) Create(ctx context.Context, conversation model.Conversation) (error, model.Conversation) {
	var cnv model.Conversation
	cnv, err := c.conversationRepo.GetByReceiverID(ctx, conversation.ReceiverID)
	if err != nil {
		return err, model.Conversation{}
	}

	if cnv.ID == 0 {
		err, cnv = c.conversationRepo.Create(ctx, conversation)
		if err != nil {
			return err, model.Conversation{}
		}
	}

	return nil, cnv
}

func (c *conversationUsecase) GetByUserID(ctx context.Context, userID uint) ([]responses.Conversation, error) {
	var cnv []responses.Conversation
	cnv, err := c.conversationRepo.GetByUserID(ctx, userID)
	if err != nil {
		return []responses.Conversation{}, err
	}

	return cnv, nil
}

func (c *conversationUsecase) GetByReceiverID(ctx context.Context, receiverID uint) (model.Conversation, error) {
	panic("implement me")
}

func InitializeConversation(conversationRepo interfaces.ConversationRepository) *conversationUsecase {
	return &conversationUsecase{conversationRepo: conversationRepo}
}
