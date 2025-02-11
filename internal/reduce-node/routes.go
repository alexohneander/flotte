package reducenode

import "github.com/gin-gonic/gin"

func registerRoutes(router *gin.Engine) *gin.Engine {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
