package role

import (
	accountusecase "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/role/usecase"

	// "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities/token"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "github.com/satori/go.uuid"
	// "github.com/google/uuid"
	"net/http"
	// "reflect"
	"strconv"
	"strings"
)

type Role struct {
	RoleUsecase    usecase.IRoleUsecase
	AccountUsecase accountusecase.IAccountUsecase
}

func (a Role) Role(r *gin.RouterGroup) {
	r.GET(utilities.GET_ROLE, a.GetRole)        //Hanya untuk Admin
	r.POST(utilities.CREATE_ROLE, a.CreateRole) //Hanya untuk Admin
	r.PUT(utilities.UPDATE_ROLE, a.UpdateRole)  //Hanya untuk Admin
}

func (a Role) GetRole(c *gin.Context) {
	tokenstring := token.ExtractToken(c)
	metadata, _ := token.ExtractTokenMetadata(tokenstring)
	email := metadata.Email
	isAdmin := a.AccountUsecase.CheckUserIsAdmin(email)
	if isAdmin {
		id, _ := c.GetQuery("id")
		orderby, _ := c.GetQuery("orderby")
		response := a.RoleUsecase.GetRoles(id, orderby)
		c.JSON(response.Status, response)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "you are not allowed",
	})
}

func (a Role) CreateRole(c *gin.Context) {
	tokenstring := token.ExtractToken(c)
	metadata, _ := token.ExtractTokenMetadata(tokenstring)
	email := metadata.Email
	isAdmin := a.AccountUsecase.CheckUserIsAdmin(email)
	if isAdmin {
		name := c.PostForm("role_name")
		form_name := models.FormName{Name: strings.Trim(name, " ")}
		response := a.RoleUsecase.CreateRole(form_name)
		c.JSON(response.Status, response)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "you are not allowed",
	})
}

func (a Role) UpdateRole(c *gin.Context) {
	tokenstring := token.ExtractToken(c)
	metadata, _ := token.ExtractTokenMetadata(tokenstring)
	email := metadata.Email
	isAdmin := a.AccountUsecase.CheckUserIsAdmin(email)
	if isAdmin {
		id := c.PostForm("role_id")
		name := c.PostForm("role_name")
		idtrim := strings.Trim(id, " ")
		nametrim := strings.Trim(name, " ")
		intid, _ := strconv.Atoi(idtrim)
		role := models.Role{Id: intid, Name: nametrim}
		response := a.RoleUsecase.UpdateRole(role)
		c.JSON(response.Status, response)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "you are not allowed",
	})
}

// update
// delete
