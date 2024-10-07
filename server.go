package main

import (
	"github.com/gin-gonic/gin"
	"test_go_1/controller"
	"test_go_1/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {

	server := gin.Default()

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/posts", func(ctx *gin.Context) {

		ctx.JSON(200, videoController.Save(ctx))

	})

	server.Run(":8080")

}
