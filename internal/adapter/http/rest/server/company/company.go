package company

import (
	"errors"
	"net/http"
	appErr "template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module/company"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type CompanyHandler interface {
	CreateCompany(c *gin.Context)
	GetCompanyById(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type companyHandler struct {
	compUsecase company.Usecase
	trans       ut.Translator
}

func CompanyInit(compUsecase company.Usecase, trans ut.Translator) CompanyHandler {
	return &companyHandler{
		compUsecase,
		trans,
	}
}

func (ch companyHandler) CreateCompany(c *gin.Context) {
	var insertCompany model.Company
	if err := c.ShouldBind(&insertCompany); err != nil {
		var verr validator.ValidationErrors

		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": verr.Translate(ch.trans)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": appErr.NewErrorResponse(appErr.ErrUnknown)})
		return

	}
	company, err := ch.compUsecase.CreateCompany(&insertCompany)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": appErr.NewErrorResponse(appErr.ErrUnknown)})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"company": company})
	return
}

func (ch companyHandler) GetCompanyById(c *gin.Context) {

}

func (ch companyHandler) DeleteUser(c *gin.Context) {}
