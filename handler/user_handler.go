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

type UserRegister struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// HandleRegisterUser godoc
// @Summary      Create new user
// @Description  Create new user and get a token
// @Produce      json
// @Accept       json
// @Param user body UserRegister true "User"
// @Success      200  {object}  ent.User
// @Success      400  {object}  model.Response
// @Router       /register [post]
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

	token, err := auth.GenerateJWT(user.ID)
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

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// HandleLoginUser godoc
// @Summary      Login
// @Description  Login to existing user and get a token
// @Produce      json
// @Accept       json
// @Param user body UserLogin true "User"
// @Success      200  {object}  ent.User
// @Success      400  {object}  model.Response
// @Router       /login [post]
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

	token, err := auth.GenerateJWT(user.ID)
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
