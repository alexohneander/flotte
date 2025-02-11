package mapnode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alexohneander/flotte/pkg/types/request"
	"github.com/gin-gonic/gin"
)

func Start() {
	err := registerAsMapNode()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router = registerRoutes(router)

	router.Run("localhost:4000") // listen and serve on 0.0.0.0:4000
}

func registerAsMapNode() error {
	// http post request to general-node
	serviceRegisterReq := request.ServiceRegister{
		Name:     "map-01",
		NodeType: "map",
		Address:  "localhost",
		Port:     "4000",
	}

	json_data, err := json.Marshal(serviceRegisterReq)
	if err != nil {
		log.Fatal(err)
		return err
	}

	resp, err := http.Post("http://localhost:8000/service", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("failed to register as map-node")
		return err
	}

	defer resp.Body.Close()
	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res)

	return nil
}
