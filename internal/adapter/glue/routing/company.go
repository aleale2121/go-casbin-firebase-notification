package routing

import (
	"template/internal/adapter/http/rest/server/company"

	"github.com/gin-gonic/gin"
)

// CompanyRoutes registers users companies
func CompanyRoutes(grp *gin.RouterGroup, compHandler company.CompanyHandler) {
	grp.POST("/companies", compHandler.CreateCompany)
	// grp.GET("/users", usrHandler.GetUsers)
	// grp.GET("/users/:id", usrHandler.GetUserById)
	// grp.DELETE("/users/:id", usrHandler.DeleteUser)
}
