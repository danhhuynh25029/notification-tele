package http

import "github.com/gin-gonic/gin"

func (u *UserHandler) NewUserRoute(r *gin.RouterGroup) {
	r.GET("/users", u.GetUser)
	r.POST("/users", u.AddUser)
	r.PUT("/users", u.UpdateUser)
}
