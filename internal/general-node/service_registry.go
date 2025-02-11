package generalnode

import (
	"fmt"

	"github.com/alexohneander/flotte/pkg/models"
	"github.com/alexohneander/flotte/pkg/types/request"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func registerSerivce(req request.ServiceRegister) error {
	db, err := gorm.Open(sqlite.Open("flotte.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
		return err
	}

	db.Create(&models.Service{Name: req.Name, NodeType: req.NodeType, Address: req.Address, Port: req.Port, Status: "running"})

	return nil
}
