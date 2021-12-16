package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/matias-ziliotto/test-golang/internal/user"
	"github.com/matias-ziliotto/test-golang/pkg/web"
)

type User struct {
	userService user.Service
}

func NewUser(userService user.Service) *User {
	return &User{
		userService: userService,
	}
}

func (u *User) GetAll(ctx *gin.Context) {
	users, err := u.userService.GetAll(ctx)

	if err != nil {
		web.Error(ctx, 500, err.Error())
		return
	}

	web.Success(ctx, 200, users)
}
