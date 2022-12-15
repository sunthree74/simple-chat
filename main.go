package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sunthree74/simple-chat/config"
	"github.com/sunthree74/simple-chat/handler"
	"github.com/sunthree74/simple-chat/repository"
	"github.com/sunthree74/simple-chat/usecase"
	"log"
	"time"
)

func main() {
	db := config.Connect()
	gin.SetMode("debug")
	db = db.Debug()
	gin.DisableConsoleColor()
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Methods",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Authorization",
			"Cookie",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(gin.Recovery())

	conversationRepo := repository.InitializeConversation(db)
	messageRepository := repository.InitializeMessage(db)

	conversationUsecase := usecase.InitializeConversation(conversationRepo)
	messageUsecase := usecase.InitializeMessage(messageRepository, conversationRepo)

	conversationHandler := handler.HandleConversation(conversationUsecase, messageUsecase)
	messageHandler := handler.HandleMessage(messageUsecase)

	router.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("simple chat service"))
		return
	})

	chatRoute := router.Group("/chat")
	{
		chatRoute.POST("/start", conversationHandler.StartConversation())
		chatRoute.POST("/send", messageHandler.SendMessage())
		chatRoute.GET("/conversation/:id", messageHandler.GetMessageByConversationId())
	}

	conversationRoute := router.Group("/conversation")
	{
		conversationRoute.GET("/me/:id", conversationHandler.GetConversationByUser())
		conversationRoute.GET("/read/:id", conversationHandler.ReadConversation())
	}

	log.Fatalln(router.Run(":8008"))
}
