package main

import (
	_ "github.com/mayur-tolexo/demo/user/docs"
	"github.com/mayur-tolexo/demo/user/handler"

	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title User API
// @description User microservice server.
// @schemes http https
// @securityDefinitions.basic BasicAuth
func main() {
	r := gin.Default()

	r.GET("/users/", handler.ListUser())
	r.POST("/user/", handler.CreateUser())
	r.GET("/swagger/*any", func(c *gin.Context) {
		httpSwagger.WrapHandler(c.Writer, c.Request)
	})
	r.Run()
}
