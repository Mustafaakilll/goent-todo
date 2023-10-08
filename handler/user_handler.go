package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustafaakilll/ent_todo/auth"
	"github.com/mustafaakilll/ent_todo/ent"
	"github.com/mustafaakilll/ent_todo/repository"
)

type UserHandler struct {
	UserRepository repository.UserRepository
}

func NewUserHandler(repository repository.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: repository,
	}
}

func (uh UserHandler) HandleRegisterUser(ctx *gin.Context) {

	var user ent.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := uh.UserRepository.RegisterUser(&user)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.GenerateJWT(user.Email, user.Fullname, user.ID)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":  newUser,
		"token": token,
	})
}

func (uh UserHandler) HandleLoginUser(ctx *gin.Context) {
	var incomingUser *ent.User
	if err := ctx.ShouldBind(&incomingUser); err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.UserRepository.GetUserByEmail(incomingUser.Email)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = repository.CheckPassword(incomingUser.Password, user.Password)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.GenerateJWT(user.Email, user.Fullname, user.ID)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})

}
