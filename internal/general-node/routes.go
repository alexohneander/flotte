package generalnode

import (
	"encoding/json"
	"fmt"

	"github.com/alexohneander/flotte/pkg/models"
	"github.com/alexohneander/flotte/pkg/types/request"
	respone "github.com/alexohneander/flotte/pkg/types/response"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	// Start Transcoding
	router.GET("/transcode", func(c *gin.Context) {
		db, err := gorm.Open(sqlite.Open("flotte.db"), &gorm.Config{})
		if err != nil {
			fmt.Println("failed to connect database")
		}

		// Get file from tmp folder
		moviePath := "./tmp/bbb_sunflower_1080p_60fps_normal.mp4"

		// Get map node counter
		var mapNodes []models.Service
		db.Where("node_type = ?", "map").Find(&mapNodes)

		// Cut Video in 10sek Segments
		i := 0
		iframes := 0
		for i < 10 {
			segmentPath := fmt.Sprintf("./tmp/seg_%d.mp4", i)
			err := ffmpeg.Input(moviePath, ffmpeg.KwArgs{"ss": iframes}).
				Output(segmentPath, ffmpeg.KwArgs{"t": 10}).
				OverWriteOutput().
				Run()

			if err != nil {
				fmt.Println(err)
			}

			i++
			iframes = iframes + 10
		}

		// send n segments to n map-nodes
		fmt.Println(mapNodes)

		c.JSON(200, gin.H{"status": "start transcoding"})
	})
	return router
}
