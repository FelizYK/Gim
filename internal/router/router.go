package router

import (
	"Gim/docs"
	"Gim/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/index", service.GetIndex)

	r.POST("/user/createUser", service.CreateUser)
	r.GET("/user/getUserList", service.GetUserList)
	r.PUT("/user/updateUser", service.UpdateUser)
	r.DELETE("/user/deleteUser", service.DeleteUser)

	return r
}
