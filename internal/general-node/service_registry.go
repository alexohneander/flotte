package generalnode

import (
	"fmt"
	"net/http"
	"time"

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

	err = db.Create(&models.Service{Name: req.Name, NodeType: req.NodeType, Address: req.Address, Port: req.Port, Status: "running"}).Error
	if err != nil {
		fmt.Println("failed to create service")
		return err
	}

	return nil
}

func checkHealth() {
	db, err := gorm.Open(sqlite.Open("flotte.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}

	for {
		var services []models.Service
		db.Find(&services, models.Service{})
		for _, service := range services {
			resp, err := http.Get("http://" + service.Address + ":" + service.Port + "/ping")
			if err != nil {
				fmt.Println("error: ", err)
				fmt.Println("delete service ", service.Name)
				db.Unscoped().Delete(&service, service.ID)
			}

			if resp != nil && resp.StatusCode != 200 {
				fmt.Println("delete service ", service.Name)
				db.Unscoped().Delete(&service, service.ID)
			}

			if resp != nil && resp.StatusCode == 200 {
				fmt.Println("service ", service.Name, " is healthy")
			}
		}

		time.Sleep(30 * time.Second)
	}

}
