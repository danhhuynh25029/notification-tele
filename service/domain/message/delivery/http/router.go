package http

import "github.com/gin-gonic/gin"

func (m *MessageHandler) NewMessageRoute(r *gin.RouterGroup) {
	r.GET("/messages", m.GetAllMessage)
	r.POST("/messages", m.AddMessage)
	r.PUT("/messages", m.UpdateMessage)
	r.DELETE("/messages", m.DeleteMessage)
}
