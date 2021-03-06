package routing

import (
	"template/internal/adapter/http/rest/server/role"

	"github.com/gin-gonic/gin"
)

// UserRoutes registers users routes
func RoleRoutes(grp *gin.RouterGroup, roleHandler role.RolesHandler) {
	roleGrp := grp.Group("/roles")
	roleGrp.POST("", roleHandler.MiddleWareValidateRole, roleHandler.AddRole)
	roleGrp.GET("/:name", roleHandler.GetRoleByName)
	roleGrp.DELETE("/:name", roleHandler.DeleteRole)
	roleGrp.GET("", roleHandler.GetRoles)
}
