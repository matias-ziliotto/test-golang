package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/matias-ziliotto/test-golang/cmd/server/validations"
	"github.com/matias-ziliotto/test-golang/internal/user"
	"github.com/matias-ziliotto/test-golang/pkg/web"
)

type storeUserRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	DocumentTypeId int    `json:"document_type_id"`
	DocumentNumber int    `json:"document_number"`
}

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
		web.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	web.Success(ctx, http.StatusOK, users)
}

func (u *User) Get(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		web.Error(ctx, http.StatusBadRequest, err.Error())
	}

	user, err := u.userService.Get(ctx, userId)

	if err != nil {
		web.Error(ctx, http.StatusNotFound, err.Error())
		return
	}

	web.Success(ctx, http.StatusOK, user)
}

func (u *User) Store(ctx *gin.Context) {
	var storeUserRequest storeUserRequest
	err := ctx.ShouldBindJSON(&storeUserRequest)

	if err != nil {
		web.Error(ctx, http.StatusBadRequest, err.Error())
	}

	validated, errMessage := validation.ValidateRequiredData(storeUserRequest, []string{"FirstName", "LastName", "DocumentTypeId", "DocumentNumber"})

	if !validated {
		web.Error(ctx, http.StatusBadRequest, errMessage)
		return
	}

	user, err := u.userService.Store(ctx, storeUserRequest.FirstName, storeUserRequest.LastName, storeUserRequest.DocumentTypeId, storeUserRequest.DocumentNumber)

	if err != nil {
		web.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	web.Success(ctx, http.StatusCreated, user)
}
