package user

import (
	"errors"
	"net/http"
	errModel "template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"

	ut "github.com/go-playground/universal-translator"
)

// UserHandler contans a function of handlers for the domian file
type UserHandler interface {
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

//UserInit initializes a user handler for the domin user
func UserInit(userUsecase user.Usecase, trans ut.Translator) UserHandler {
	return &userHandler{
		userUsecase,
		trans,
	}
}

// CreateUser creates a new user
// POST /v1/:com-id/users
func (uh userHandler) CreateUser(c *gin.Context) {
	ID := c.Param("comp-id")

	compID, err := uuid.FromString(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errModel.NewErrorResponse(err)})
		return
	}

	var insertUser model.User

	if err := c.ShouldBind(&insertUser); err != nil {

		var verr validator.ValidationErrors

		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": verr.Translate(uh.trans)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errModel.NewErrorResponse(errModel.ErrUnknown)})
		return

	}
	user, err := uh.userUsecase.CreateUser(compID, &insertUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errModel.NewErrorResponse(err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
	return
}

// GetUserById gets a user by id
func (uh userHandler) GetUserById(c *gin.Context) {

	ID := c.Param("id")

	id, err := uuid.FromString(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errModel.NewErrorResponse(err)})
		return
	}

	user, err := uh.userUsecase.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": errModel.NewErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
	return
}

// DeleteUser deletes user by id
func (uh userHandler) DeleteUser(c *gin.Context) {
	ID := c.Param("id")

	id, err := uuid.FromString(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errModel.NewErrorResponse(err)})
		return
	}

	err = uh.userUsecase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": errModel.NewErrorResponse(err)})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
	return
}

// GetUsers gets a list of users
func (uh userHandler) GetUsers(c *gin.Context) {
	users, err := uh.userUsecase.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": errModel.NewErrorResponse(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
	return
}
