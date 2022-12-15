package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sunthree74/simple-chat/interfaces"
	"github.com/sunthree74/simple-chat/model"
	"github.com/sunthree74/simple-chat/structs/requests"
	"github.com/sunthree74/simple-chat/structs/responses"
	"net/http"
	"strconv"
)

type messageHandler struct {
	messageUsecase interfaces.MessageUsecase
}

func HandleMessage(messageUsecase interfaces.MessageUsecase) *messageHandler {
	return &messageHandler{messageUsecase: messageUsecase}
}

func (c *messageHandler) SendMessage() gin.HandlerFunc {
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

		var msg model.Message
		msg.Text = form.Text
		msg.ConversationID = form.ConversationID
		msg.MessageReplyID = form.MessageReplyID
		if err := c.messageUsecase.Create(ctx, msg); err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "status": http.StatusBadRequest, "message": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"success": true, "status": http.StatusOK, "message": "Pesan telah terkirim"})
		return
	}
}

func (c *messageHandler) GetMessageByConversationId() gin.HandlerFunc {
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

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": nil})
			return
		}

		var messages []responses.Message
		if messages, err = c.messageUsecase.GetByConversationID(ctx, uint(id)); err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "status": http.StatusBadRequest, "message": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"success": true, "status": http.StatusOK, "data": messages})
		return
	}
}
