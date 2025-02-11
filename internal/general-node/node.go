package generalnode

import (
	"github.com/alexohneander/flotte/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Start() {
	// Init Service Registry Database
	db, err := gorm.Open(sqlite.Open("flotte.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Service{})

	// Init API Backend
	router := gin.Default()
	router = registerRoutes(router)

	router.Run("localhost:8000") // listen and serve on 0.0.0.0:8000
}
