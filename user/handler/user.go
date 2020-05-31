package handler

import (
	"net/http"

	"github.com/mayur-tolexo/demo/user/model"

	"github.com/mayur-tolexo/demo/db"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

// ListUser is the handler of list user endpoint
// @Summary List users
// @Description list all the users based on filter given
// @Tags user
// @Produce  json
// @Success 200 {object} model.UserList
// @Router /users/ [get]
func ListUser() func(c *gin.Context) {
	users := model.UserList{[]model.UserDetail{model.UserDetail{db.User{Name: "Test"}}}}
	return func(c *gin.Context) {
		c.JSON(200, users)
	}
}

// CreateUser is create user endpoint handler
// @Summary create user
// @Description create a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body model.CreateUser true "create user"
// @Success 200 {object} model.UserDetail
// @Failure 400 {object} httputil.HTTPError
// @Router /user/ [post]
func CreateUser() func(c *gin.Context) {
	user := model.UserDetail{}
	return func(c *gin.Context) {
		var req model.CreateUser
		if err := c.ShouldBindJSON(&req); err != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
			return
		}
		user.Name = req.Name
		c.JSON(200, user)
	}
}
