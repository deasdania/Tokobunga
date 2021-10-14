package role

import (
	authusecase "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/auth/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/role/usecase"

	// "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "github.com/satori/go.uuid"
	// "github.com/google/uuid"
	// "net/http"
	// "reflect"
)

type Role struct {
	RoleUsecase usecase.IRoleUsecase
	AuthUsecase authusecase.IAuthUsecase
}

func (a Role) Role(r *gin.RouterGroup) {
	r.GET(utilities.GET_ROLE, a.GetRole)
	r.POST(utilities.CREATE_ROLE, a.CreateRole)
}

func (a Role) GetRole(c *gin.Context) {
	id, _ := c.GetQuery("id")
	orderby, _ := c.GetQuery("orderby")
	response := a.RoleUsecase.GetRoles(id, orderby)
	c.JSON(response.Status, response)
}

func (a Role) CreateRole(c *gin.Context) {
	name := c.PostForm("role_name")
	form_name := models.FormName{Name: name}
	response := a.RoleUsecase.CreateRole(form_name)
	c.JSON(response.Status, response)
}
