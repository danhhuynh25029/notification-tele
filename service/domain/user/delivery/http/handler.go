package http

import (
	"bot/service/domain/user/usecase"
	"bot/service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userUseCase usecase.IUserUsecase
}

func NewUserHandler(userUseCase usecase.IUserUsecase) UserHandler {
	return UserHandler{
		userUseCase: userUseCase,
	}
}

func (u *UserHandler) AddUser(c *gin.Context) {
	var (
		userRequest models.User
	)
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot parse json into object",
		})
		return
	}
	err = u.userUseCase.AddUser(c, userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot add user",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
	})
}
func (u *UserHandler) UpdateUser(c *gin.Context) {
	var (
		userRequest models.User
	)
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot parse json into object",
		})
		return
	}
	err = u.userUseCase.UpdateUser(c, userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot update user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
	})
}
func (u *UserHandler) GetUser(c *gin.Context) {
	var req models.User
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot get user",
		})
		return
	}
	users, err := u.userUseCase.GetUser(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot get user",
		})
		return
	}
	c.JSON(http.StatusOK, users)
}
