package http

import (
	"bot/service/domain/message/usecase"
	"github.com/gin-gonic/gin"
)

type IMessageHandler interface {
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

}
func (m *MessageHandler) UpdateMessage(c *gin.Context) {

}
func (m *MessageHandler) DeleteMessage(c *gin.Context) {

}

func (m *MessageHandler) GetAllMessage(c *gin.Context) {

}
