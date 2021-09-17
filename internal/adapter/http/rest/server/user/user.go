package user

import (
	"errors"
	"log"
	"net/http"
	"template/internal/constant/model"
	"template/internal/module/user"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// UserHandler contans a function of handlers for the domian file
type UsersHandler interface {
	CreateUser(c *gin.Context)
	GetUserById(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUsers(c *gin.Context)
}

// userHandler defines all the things neccessary for users handlers
type userHandler struct {
	userUsecase user.Usecase
	trans       ut.Translator
}

//UserInit initializes a user handler for the domin file
func UserInit(userUsecase user.Usecase, trans ut.Translator) UsersHandler {
	return &userHandler{
		userUsecase,
		trans,
	}
}

// CreateUser creates a new user
// POST /v1/users
func (uh userHandler) CreateUser(c *gin.Context) {
	var insertUser model.User

	if err := c.ShouldBind(&insertUser); err != nil {
		
		var verr validator.ValidationErrors

		if errors.As(err, &verr) {
			log.Println(err)
			log.Println(err)
			log.Println(err)
			log.Println(err)
			log.Println(err)
			log.Println(err)
			log.Println(err)
			log.Println(err)
			log.Println(err)
			c.JSON(http.StatusBadRequest, verr.Translate(uh.trans))
			return
		}
	}

	user, err := uh.userUsecase.CreateUser(&insertUser)

	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
	return
}

// GetUserById gets a user by id
func (uh userHandler) GetUserById(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{})
	return
}

// DeleteUser deletes user by id
func (uh userHandler) DeleteUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{})
	return
}

// GetUsers gets a list of users
func (uh userHandler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{})
	return
}
