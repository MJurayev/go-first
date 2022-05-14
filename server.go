package main

import (
	"github.com/gin-gonic/gin"
	"ictm.uz/video/controller"
	"ictm.uz/video/service"
)



func main(){
	videoService := service.New()
	videoController :=controller.New(videoService)
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":"OK",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		videos := videoController.FindAll()
		ctx.JSON(200, gin.H{
			"data":videos,
		})
	})

	server.POST("/videos", func(ctx *gin.Context) {
		video := videoController.Save(ctx)
		ctx.JSON(201, gin.H{
			"data":video,
		})
	})

	server.Run(":8080")
}