package main

import (
	"github.com/mayur-tolexo/demo/user/docs"
	"github.com/mayur-tolexo/demo/user/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.basic BasicAuth
func main() {
	docs.SwaggerInfo.Title = "User API"
	docs.SwaggerInfo.Description = "User microservice server."
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	r.GET("/users/", handler.ListUser())
	r.POST("/user/", handler.CreateUser())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
