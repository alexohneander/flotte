package generalnode

import (
	"encoding/json"

	"github.com/alexohneander/flotte/pkg/types/request"
	respone "github.com/alexohneander/flotte/pkg/types/response"
	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) *gin.Engine {
	// Get Status
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Add Service
	router.POST("/service", func(c *gin.Context) {
		defer c.Request.Body.Close()
		var req request.ServiceRegister
		err := json.NewDecoder(c.Request.Body).Decode(&req)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		}

		err = registerSerivce(req)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(200, respone.ServiceRegister{
			ServiceName: req.Name,
			Status:      "registered",
		})
	})

	return router
}
