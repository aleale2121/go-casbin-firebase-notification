package routing

import (
	"template/internal/adapter/http/rest/server/role"

	"github.com/gin-gonic/gin"
)

// UserRoutes registers users routes
func RoleRoutes(grp *gin.RouterGroup, roleHandler role.RolesHandler) {
	grp.POST("",roleHandler.MiddleWareValidateRole, roleHandler.AddRole)
	grp.GET("/:name", roleHandler.GetRoleByName)
	grp.DELETE("/:name", roleHandler.DeleteRole)
	grp.GET("", roleHandler.GetRoles)
}
