package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sunthree74/simple-chat/interfaces"
	"github.com/sunthree74/simple-chat/model"
	"github.com/sunthree74/simple-chat/structs/requests"
	"net/http"
)

type conversationHandler struct {
	conversationUsecase interfaces.ConversationUsecase
	messageUsecase interfaces.MessageUsecase
}

func HandleConversation(usecase interfaces.ConversationUsecase, messageUsecase interfaces.MessageUsecase) *conversationHandler {
	return &conversationHandler{conversationUsecase: usecase, messageUsecase: messageUsecase}
}

func (c *conversationHandler) StartConversation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//claimVal, isExists := ctx.Get("UserID")
		//if !isExists {
		//	ctx.JSON(http.StatusUnauthorized, gin.H{"success": false, "status": http.StatusUnauthorized, "message": "Must login or Invalid Token."})
		//	return
		//}
		//
		//userID, err := strconv.ParseUint(claimVal.(*model.Claim).UserID, 10, 32)
		//if err != nil {
		//	ctx.JSON(http.StatusUnauthorized, gin.H{"success": false, "status": http.StatusUnauthorized, "message": "Must login or Invalid Token."})
		//	return
		//}

		var form requests.ChatParam
		if err := ctx.ShouldBindJSON(&form); err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "status": http.StatusBadRequest, "message": "Invalid parameter."})
			return
		}

		if form.Text == ""{
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "status": http.StatusBadRequest, "message": "Text must not be empty."})
			return
		}

		var conversation model.Conversation
		conversation.ReceiverID = form.ReceiverID
		conversation.UserID = form.ReceiverID
		err,  conversation := c.conversationUsecase.Create(ctx, conversation);
		if err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "status": http.StatusBadRequest, "message": err})
			return
		}

		var msg model.Message
		msg.Text = form.Text
		msg.ConversationID = conversation.ID
		if err := c.messageUsecase.Create(ctx, msg); err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "status": http.StatusBadRequest, "message": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"success": true, "status": http.StatusOK, "message": "Pesan telah terkirim"})
		return
	}
}
