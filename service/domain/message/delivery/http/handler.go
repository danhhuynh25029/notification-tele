package http

import (
	"bot/service/domain/message/usecase"
	"bot/service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IMessageHandler interface {
	NewMessageRoute(r *gin.RouterGroup)
}
type MessageHandler struct {
	messageUseCase usecase.IMessageUsecase
}

func NewMessageHandler(messageUseCase usecase.IMessageUsecase) IMessageHandler {
	return &MessageHandler{
		messageUseCase: messageUseCase,
	}
}
func (m *MessageHandler) AddMessage(c *gin.Context) {
	var message models.Message
	err := c.ShouldBindJSON(&message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}
	err = m.messageUseCase.AddMessage(c, message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
	})
}
func (m *MessageHandler) UpdateMessage(c *gin.Context) {
	var message models.Message
	err := c.ShouldBindJSON(&message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}
	err = m.messageUseCase.UpdateMesasge(c, message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
	})
}
func (m *MessageHandler) DeleteMessage(c *gin.Context) {
	var message models.Message
	err := c.ShouldBindJSON(&message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}
	err = m.messageUseCase.DeleteMessage(c, message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
	})
}

func (m *MessageHandler) GetAllMessage(c *gin.Context) {
	var (
		messages []models.Message
		err      error
	)
	messages, err = m.messageUseCase.GetAllMessage(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}
	c.JSON(http.StatusOK, messages)
}
