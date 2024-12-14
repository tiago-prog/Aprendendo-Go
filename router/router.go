package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})

	})

	router.Run(":3035")
}
