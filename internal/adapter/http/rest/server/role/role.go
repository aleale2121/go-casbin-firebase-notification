package role

import (
	"template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module/role"

	"github.com/gin-gonic/gin"
)

type RolesHandler interface {
	MiddleWareValidateRole(c *gin.Context)
	GetRoles(c *gin.Context)
	GetRoleByName(c *gin.Context)
	AddRole(c *gin.Context)
	DeleteRole(c *gin.Context)
}
type rolesHandler struct {
	roleUseCase        role.UseCase
}

func NewRoleHandler(useCase role.UseCase) RolesHandler {
	return &rolesHandler{roleUseCase: useCase}
}
func (n rolesHandler) MiddleWareValidateRole(c *gin.Context) {
	roleX := model.Role{}
	err := c.Bind(&roleX)
	if err != nil {
		c.JSON(422,err)
		return
	}
	c.Set("x-role", roleX)
	c.Next()
}

func (n rolesHandler) GetRoles(c *gin.Context) {
	roles, err := n.roleUseCase.Roles()

	if err != nil {
		c.JSON(errors.GetStatusCode(err),err)
		return
	}
	c.JSON(200,roles)

}

func (n rolesHandler) GetRoleByName(c *gin.Context) {
	rolename := c.Param("name")

	r, err := n.roleUseCase.Role(rolename)

	if err != nil {
		c.JSON(errors.GetStatusCode(err),err)
		return
	}
	c.JSON(200,r)

}

func (n rolesHandler) AddRole(c *gin.Context) {
	rl := c.MustGet("x-role").(model.Role)

	r, err := n.roleUseCase.StoreRole(rl)

	if err != nil {
				c.JSON(errors.GetStatusCode(err),err)

		return
	}

		c.JSON(200,r)


}

func (n rolesHandler) DeleteRole(c *gin.Context) {

	rolename := c.Param("name")
	err := n.roleUseCase.DeleteRole(rolename)

	if err != nil {
				c.JSON(errors.GetStatusCode(err),err)

		return
	}
	c.JSON(200,"User Delted Successfully")
}
