package generalnode

import "github.com/gin-gonic/gin"

func Start() {
	router := gin.Default()

	// Get Status
	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Add Service
	router.POST("/service", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"serviceName": "service1",
			"status":      "registered",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
