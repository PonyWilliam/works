package server

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine{
	ginRouter := gin.Default()
	ginRouter.POST("/test/",func(context *gin.Context){
		context.String(200,"test")
	})
	return ginRouter
}